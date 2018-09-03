package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type whoiam struct {
	Addr string
}

func main() {

	url := ""

	if "" != os.Getenv("BROKER_URL") {
		url = os.Getenv("BROKER_URL")
	}

	username := ""
	password := ""
	endpoint := ""

	if "" != os.Getenv("BROKER_USERNAME") {
		username = os.Getenv("BROKER_USERNAME")
	}
	if "" != os.Getenv("BROKER_PASSWORD") {
		password = os.Getenv("BROKER_PASSWORD")
	}
	if "" != os.Getenv("BROKER_ENDPOINT") {
		endpoint = os.Getenv("BROKER_ENDPOINT")
	}

	basic := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))
	log.Printf("Target %s.", url)

	req, err := http.NewRequest("GET", url+endpoint, nil)
	req.Header.Add("Authorization", basic)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	println("You are " + string(body))
}
