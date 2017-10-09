package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	b      = "\xF0\x9F\x85\xB1"
	a      = "\xF0\x9F\x85\xB0"
	o      = "\xF0\x9F\x85\xBE"
	p      = "\xF0\x9F\x85\xBF"
	ab     = "\xF0\x9F\x86\x8E"
	cl     = "\xF0\x9F\x86\x91"
	c      = "\xC2\xA9"
	r      = "\xC2\xAE"
	i      = "\xE2\x84\xB9"
	x      = "\xE2\x9D\x8C"
	qmark  = "\xE2\x9D\x93"
	emark  = "\xE2\x9D\x97"
	docker = "\xF0\x9F\x90\xB3"
	one    = "\x31\xE2\x83\xA3"
	two    = "\x32\xE2\x83\xA3"
	three  = "\x33\xE2\x83\xA3"
	four   = "\x34\xE2\x83\xA3"
	five   = "\x35\xE2\x83\xA3"
	six    = "\x36\xE2\x83\xA3"
	seven  = "\x37\xE2\x83\xA3"
	eight  = "\x38\xE2\x83\xA3"
	nine   = "\x39\xE2\x83\xA3"
	zero   = "\x30\xE2\x83\xA3"
	dexcl  = "\xE2\x80\xBC"
	exclq  = "\xE2\x81\x89"
	beer   = "\xF0\x9F\x8D\xBA"
	dbeer  = "\xF0\x9F\x8D\xBB"
	terr   = "\xF0\x9F\x91\xB3 \xF0\x9F\x92\xA3"
	ngul   = "\xF0\x9F\x86\x96 \xF0\x9F\x86\x92"
	suh    = "suh my ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
	sos    = "\xF0\x9F\x86\x98"
	ok     = "\xF0\x9F\x86\x97"
)

var (
	replacer strings.Replacer
	apiKey   string
	certPath string
	keyPath  string
	debug    bool
	domain   string
	port     string
)

func main() {
	setupReplacer()
	setupParams()

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = debug

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.RemoveWebhook()
	if err != nil {
		log.Fatal(err)
	}
	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://" + domain + ":" + port + "/" + apiKey))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + apiKey)
	go http.ListenAndServeTLS("0.0.0.0:"+port, certPath, keyPath, nil)
	for update := range updates {
		if update.InlineQuery != nil {
			log.Printf("[INLINE] new query sent in by %s -> %s\n", update.InlineQuery.From.UserName, update.InlineQuery.Query)
			memeStr := nibbering(strings.ToUpper(update.InlineQuery.Query))

			article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, suh, memeStr)
			article.Description = memeStr

			inlineConf := tgbotapi.InlineConfig{
				InlineQueryID: update.InlineQuery.ID,
				IsPersonal:    true,
				CacheTime:     0,
				Results:       []interface{}{article},
			}

			if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
				log.Println(err)
			}
		} else if update.Message != nil {
			log.Printf("[MESSAGE] new message sent in by %s -> %s\n", update.Message.From.UserName, update.Message.Text)
			memeStr := nibbering(strings.ToUpper(update.Message.Text))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, memeStr)
			msg.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
		}
	}

}

func nibbering(str string) string {
	return replacer.Replace(str)
}

func setupReplacer() {
	// Combo emojis (like AB, CL) goes first, otherwise
	// the Replacer completely ignores them.
	//
	// Will rewrite this shitty thingy later.
	replacer = *strings.NewReplacer(
		"2PERONI", dbeer,
		"PERONI", beer,
		"SOS", sos,
		"AB", ab,
		"CL", cl,
		"NGUL", ngul,
		"ALLAH", terr,
		"OK", ok,
		"A", a,
		"C", c,
		"G", b,
		"B", b,
		"O", o,
		"P", p,
		"R", r,
		"I", i,
		"X", x,
		"DOCKER", docker,
		"!?", exclq,
		"!!", dexcl,
		"?", qmark,
		"!", emark,
		"1", one,
		"2", two,
		"3", three,
		"4", four,
		"5", five,
		"6", six,
		"7", seven,
		"8", eight,
		"9", nine,
		"0", zero,
	)
}

func setupParams() {
	flag.StringVar(&certPath, "cert", "", "required, TLS certificate path")
	flag.StringVar(&keyPath, "key", "", "required, TLS key path")
	flag.StringVar(&apiKey, "apikey", "", "required, Telegram bot API key")
	flag.StringVar(&domain, "domain", "", "required, domain associated to the TLS cert+key and the server where this bot will be running")
	flag.StringVar(&port, "port", "88", "port to run on, must be 443, 80, 88, 8443")
	flag.BoolVar(&debug, "debug", false, "debug Telegram bot interactions")
	flag.Parse()

	if certPath == "" || keyPath == "" || apiKey == "" || domain == "" {
		flag.Usage()
		os.Exit(1)
	}
}
