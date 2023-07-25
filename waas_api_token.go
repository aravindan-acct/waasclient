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
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
	var token_map map[string]interface{}
	json.Unmarshal(body, &token_map)
	log.Println(token_map)
	api_token := token_map["key"].(string)
	return api_token
}

func WAFToken(username string, password string, WAF_IP string, Cuda_WAF_Client *http.Client) string {
	var waf_login_url string
	waf_login_url = WAF_IP + "/restapi/v3.1/login"
	Cuda_WAF_Client = WaasClient()
	var login_data map[string]string
	login_data = make(map[string]string)
	login_data["username"] = username
	login_data["password"] = password

	json_login_data, err := json.Marshal(login_data)
	log.Println("JSON Payload:")
	log.Println(string(json_login_data))
	login_payload := bytes.NewBuffer(json_login_data)
	login_req, err := http.NewRequest("POST", waf_login_url, io.Reader(login_payload))
	if err != nil {
		log.Println(err)
	}
	login_req.Header.Set("Content-Type", "application/json")

	res, err := Cuda_WAF_Client.Do(login_req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
	var token_map map[string]interface{}
	json.Unmarshal([]byte(body), &token_map)
	var token string
	token = token_map["token"].(string)
	log.Println(token + ":")
	return token + ":"
}
