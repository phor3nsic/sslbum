package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var banner = `

█▀ █▀ █░░ █▄▄ █░█ █▀▄▀█
▄█ ▄█ █▄▄ █▄█ █▄█ █░▀░█

░░░░░░███████ ]▄▄▄▄▄▄▄▄
▂▄▅█████████▅▄▃▂
I███████████████████].
◥⊙▲⊙▲⊙▲⊙▲⊙▲⊙▲⊙◤...

		by @phor3nsic

[!] Run: go run sslbum.go example.com
[!] Need set enviroment SHODAN with API key
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
		fmt.Println(concat)
		for i := 0; i < len(responseObj.Matches); i++ {
			fmt.Println(responseObj.Matches[i].IPStr)
		}

	}
}

func shodan(domain string) {
	shodanKey := os.Getenv("SHODAN")
	shodanQuery := []string{`ssl:`, "ssl.cert.subject.CN:"}
	for _, q := range shodanQuery {
		shodanApi := "https://api.shodan.io/shodan/host/search?key=" + shodanKey + "&query=" + q + domain
		request(shodanApi)
	}

}

func main() {
	fmt.Println(banner)
	fmt.Println("[-] Searching ...")
	domain := os.Args[1]
	shodan(domain)
}
