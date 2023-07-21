package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	waasclient "waasclient/waas_client"
)

// Take environment variables as input for credentials and makes an API call to WAAS.
func main() {
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
