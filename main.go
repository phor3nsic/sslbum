package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	flag.BoolVar(&silent, "s", false, "Silent Mode")
	flag.StringVar(&domain, "d", "", "Domaion to check")
	flag.StringVar(&apiKey, "k", "", "Shodan ApiKey")
	flag.Parse()
	if silent != true {
		fmt.Println(banner)
	}
	ssl(domain, apiKey)
}

var domain string
var silent bool
var apiKey string
var results []string
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

func dupCheck(val string, list []string) bool {

	for _, b := range list {
		if b == val {
			return false
		}
	}
	return true
}

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

		if silent != true {
			concat := fmt.Sprint("[+] Total found: ", responseObj.Total)
			fmt.Println(concat)
		}

		for i := 0; i < len(responseObj.Matches); i++ {
			var ip string
			ip = responseObj.Matches[i].IPStr
			if dupCheck(ip, results) {
				results = append(results, ip)
			}

		}

	}
}

func ssl(domain string, api string) {
	shodanQuery := []string{"ssl:", "ssl.cert.subject.CN:"}
	for _, q := range shodanQuery {
		shodanApi := "https://api.shodan.io/shodan/host/search?key=" + api + "&query=" + q + domain
		if silent != true {
			fmt.Println("[-] " + q + " Searching ...")
		}
		request(shodanApi)
	}

	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}

}
