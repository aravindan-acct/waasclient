package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	waasclient "waasclient/waas_client"
)

// Take environment variables as input for credentials and makes an API call to WAAS.
func waas_api() {
	email := os.Getenv("WAAS_EMAIL")
	password := os.Getenv("WAAS_PASSWD")
	httpClient := waasclient.WaasClient()
	token := waasclient.Token(email, password, httpClient)

	req, err := http.NewRequest("GET", waasclient.URL+"applications/", nil)

	if err != nil {
		log.Println(err)

	}
	req.Header.Set("auth-api", token)

	res, err := httpClient.Do(req)

	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
	log.Println(res.StatusCode)
}

func waf_api() {
	username := os.Getenv("WAF_USERNAME")
	password := os.Getenv("WAF_PASSWD")
	waf_info := os.Getenv("WAF_INFO")
	WAF_Client := waasclient.WaasClient()
	waf_token := waasclient.WAFToken(username, password, waf_info, WAF_Client)
	log.Println(waf_token)
}

func main() {
	if os.Args[1] == "waf" {
		waf_api()
	}
	if os.Args[1] == "waas" {
		waas_api()
	}
}
