package waasclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetches the api token string to make authenticated api call to waf-as-a-service.
// To make API calls, provide this token as the value for the auth-api HTTP header.

func Token(username string, password string, CudaClient *http.Client) string {
	CudaClient = WaasClient()

	payload := "email=" + username + "&password=" + password + "&account_id=9714266"
	payload_buf := bytes.NewBuffer([]byte(payload))
	req, err := http.NewRequest("POST", URL+"api_login/", io.Reader(payload_buf))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := CudaClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(body))
	var token_map map[string]interface{}
	json.Unmarshal(body, &token_map)
	log.Println(token_map)
	api_token := token_map["key"].(string)
	return api_token
}
