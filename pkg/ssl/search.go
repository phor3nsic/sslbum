package ssl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var results []string

type Response struct {
	Matches []struct {
		IPStr string `json:"ip_str"`
	} `json:"matches"`
	Total int `json:"total"`
}

func dupCheck(val string, list []string) bool {

	for _, b := range list {
		if b == val {
			return false
		}
	}
	return true
}

func Ssl(domain string, api string) {
	shodanQuery := []string{"ssl:", "ssl.cert.subject.CN:"}
	for _, q := range shodanQuery {
		shodanApi := "https://api.shodan.io/shodan/host/search?key=" + api + "&query=" + q + domain
		request(shodanApi)
	}

	for i := 0; i < len(results); i++ {
		fmt.Println(results[i])
	}

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

		for i := 0; i < len(responseObj.Matches); i++ {
			var ip string
			ip = responseObj.Matches[i].IPStr
			if dupCheck(ip, results) {
				results = append(results, ip)
			}

		}

	}
}
