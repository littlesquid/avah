package http

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient struct {
	Method        string
	Url           string
	RequestHeader map[string]string
	RequestBody   []byte
	ResponseBody  interface{}
}

func (httpClient HttpClient) request() {

	body := bytes.NewBuffer(httpClient.RequestBody)
	req, err := http.NewRequest(httpClient.Method, httpClient.Url, body)

	client := &http.Client{}
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("sending request to %v with exception %v", httpClient.Url, err)
		return
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	httpClient.ResponseBody = responseBody
}
