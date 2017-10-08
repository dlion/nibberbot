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
	b   = "\xF0\x9F\x85\xB1"
	a   = "\xF0\x9F\x85\xB0"
	o   = "\xF0\x9F\x85\xBE"
	p   = "\xF0\x9F\x85\xBF"
	ab  = "\xF0\x9F\x86\x8E"
	cl  = "\xF0\x9F\x86\x91"
	suh = "suh my ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
)

var (
	replacer strings.Replacer
	apiKey   string
	certPath string
	keyPath  string
	debug    bool
	domain   string
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
	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://" + domain + ":8443/" + apiKey))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + apiKey)
	go http.ListenAndServeTLS("0.0.0.0:8443", certPath, keyPath, nil)
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
		"AB", ab,
		"CL", cl,
		"A", a,
		"G", b,
		"B", b,
		"O", o,
		"P", p,
	)
}

func setupParams() {
	flag.StringVar(&certPath, "cert", "", "required, TLS certificate path")
	flag.StringVar(&keyPath, "key", "", "required, TLS key path")
	flag.StringVar(&apiKey, "apikey", "", "required, Telegram bot API key")
	flag.StringVar(&domain, "domain", "", "required, domain associated to the TLS cert+key and the server where this bot will be running")
	flag.BoolVar(&debug, "debug", false, "debug Telegram bot interactions")
	flag.Parse()

	if certPath == "" || keyPath == "" || apiKey == "" || domain == "" {
		flag.Usage()
		os.Exit(1)
	}
}
