package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	url := ""
	username := ""
	password := ""
	endpoint := ""

	if "" != os.Getenv("BROKER_URL") {
		url = os.Getenv("BROKER_URL")
	}
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

	// println("Response: " + string(body))

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		println("JSON parse error: ", error)
		return
	}

	println("Response: ", string(prettyJSON.Bytes()))
}
