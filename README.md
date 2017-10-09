# `nibberbot` - a Telegram bot to write like a ni🅱️🅱️a

## Usage

`nibberbot` is go get-able:

```
$ go get -u github.com/gsora/nibberbot/...
$ nibberbot
Usage of nibberbot:
  -apikey string
    	required, Telegram bot API key
  -cert string
    	required, TLS certificate path
  -debug
    	debug Telegram bot interactions
  -domain string
    	required, domain associated to the TLS cert+key and the server where this bot will be running
  -key string
    	required, TLS key path
  -port string
    	port to run on, must be 443, 80, 88, 8443 (default "88")
```

`nibberbot` works with both inline request and direct messages.

Letsencrypt users must provide `fullchain.pem` as `cert`.

## Hacking

To add more emojis, one should:

 1. add the emoji UTF-8 hex representation to the `const` declaration at line 13
 	- You can get the hex [here](https://apps.timwhitlock.info/emoji/tables/unicode)
 2. in the `setupReplacer` function, add a conversion tuple {expression to transform in emoji, emoji hex variable}

For example, if one want to add the '🅱️' emoji:

```go
const (
	b   = "\xF0\x9F\x85\xB1" // new emoji
	a   = "\xF0\x9F\x85\xB0"
	o   = "\xF0\x9F\x85\xBE"
	p   = "\xF0\x9F\x85\xBF"
	ab  = "\xF0\x9F\x86\x8E"
	cl  = "\xF0\x9F\x86\x91"
	suh = "suh my ni\xF0\x9F\x85\xB1\xF0\x9F\x85\xB1a"
)

...

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
		"B", b, // new replacer for the "B" letter, using the b emoji hex representation
		"O", o,
		"P", p,
	)
}
```

The replacement routine is clumsy and prone to error, a reworked and more maintainable solution should be created and put in place: pull requests are welcome!
