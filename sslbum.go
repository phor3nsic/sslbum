package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var domain string
var silent bool
var apiKey string
var banner = `

█▀ █▀ █░░ █▄▄ █░█ █▀▄▀█
▄█ ▄█ █▄▄ █▄█ █▄█ █░▀░█

░░░░░░███████ ]▄▄▄▄▄▄▄▄
▂▄▅█████████▅▄▃▂
I███████████████████].
◥⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙◤...

		by @phor3nsic

[!] Run: sslbum -d example.com -k SHODANAPIKEY
`

type Response struct {
	Matches []struct {
		IPStr string `json:"ip_str"`
	} `json:"matches"`
	Total int `json:"total"`
}

func request(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var responseObj Response
		json.Unmarshal([]byte(string(data)), &responseObj)
		if responseObj.Total == 0 {
			return
		}
		concat := fmt.Sprint("[+] Total found: ", responseObj.Total)
		if silent != true {
			fmt.Println(concat)
		}
		for i := 0; i < len(responseObj.Matches); i++ {
			fmt.Println(responseObj.Matches[i].IPStr)
		}

	}
}

func shodan(domain string, api string) {
	shodanQuery := []string{`ssl:`, "ssl.cert.subject.CN:"}
	for _, q := range shodanQuery {
		shodanApi := "https://api.shodan.io/shodan/host/search?key=" + api + "&query=" + q + domain
		if silent != true {
			fmt.Println("[-] Searching ...")
		}
		request(shodanApi)
	}

}

func main() {
	flag.BoolVar(&silent, "s", false, "Silent Mode")
	flag.StringVar(&domain, "d", "", "Domaion to check")
	flag.StringVar(&apiKey, "k", "", "Shodan ApiKey")
	flag.Parse()
	if silent != true {
		fmt.Println(banner)
	}
	shodan(domain, apiKey)
}
