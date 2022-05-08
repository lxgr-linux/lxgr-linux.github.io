package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RequestBody struct {
	Host        string   `json:"host"`
	Key         string   `json:"key"`
	KeyLocation string   `json:"keyLocation"`
	UrlList     []string `json:"urlList"`
}

func NewRequestBody(urlList []string) RequestBody {
	return RequestBody{
		Host:        "lxgr-linux.github.io",
		Key:         "447564c7456941a08a65db5651d3db19",
		KeyLocation: "https://lxgr-linux.github.io/447564c7456941a08a65db5651d3db19.txt",
		UrlList:     urlList,
	}
}

func sendRequest(url string, urlList []string) error {
	postBody, err := json.Marshal(NewRequestBody(urlList))
	if err != nil {
		return err
	}
	log.Printf("%s", postBody)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("Status: %s", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(body)
	log.Printf("%s", sb)
	return nil
}

func main() {
	err := sendRequest("https://yandex.com/indexnow", os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	err = sendRequest("https://bing.com/IndexNow", os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
