package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Method        string
	Url           string
	RequestHeader map[string]string
	RequestBody   []byte
	ResponseBody  string
}

func (httpClient *HttpClient) Request() {

	fmt.Println("sending request to %v", httpClient.Url)

	body := bytes.NewBuffer(httpClient.RequestBody)
	req, err := http.NewRequest(httpClient.Method, httpClient.Url, body)

	for key, value := range httpClient.RequestHeader {
		req.Header.Set(key, value)
	}

	if err != nil {
		fmt.Errorf("create request failed: ", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("sending request to %v with exception %v", httpClient.Url, err)
		return
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("read response body failed: ", err)
	}

	httpClient.ResponseBody = string(responseBody)
}
