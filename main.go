package main

import (
	"flag"
	"fmt"

	"github.com/phor3nsic/sslbum/pkg/ssl"
)

func main() {
	flag.BoolVar(&silent, "s", false, "Silent Mode")
	flag.StringVar(&target, "t", "", "Target to check")
	flag.StringVar(&apiKey, "k", "", "Shodan ApiKey")
	flag.StringVar(&mode, "mode", "ssl", "Mode to use ss/fav")
	flag.Parse()
	if silent != true {
		fmt.Println(banner)
	}
	ssl.Ssl(target, apiKey)

}

var target string
var silent bool
var apiKey string
var mode string
var banner = `

█▀ █▀ █░░ █▄▄ █░█ █▀▄▀█
▄█ ▄█ █▄▄ █▄█ █▄█ █░▀░█

░░░░░░███████ ]▄▄▄▄▄▄▄▄
▂▄▅█████████▅▄▃▂
I███████████████████].
◥⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙◤...

		by @phor3nsic

`
