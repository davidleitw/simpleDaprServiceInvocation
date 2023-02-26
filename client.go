package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	address = "http://localhost:3500"
	appID   = "echo-http-server"
)

func main() {
	for {
		log.Println("Start invocate another function")
		invokeByAppID()
		invokeByInvokeApiURL()
		invokeByDaprSDK()
		fmt.Println()
		time.Sleep(2 * time.Second)
	}
}

func invokeByAppID() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/echo?msg=hello-world", address), nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("dapr-app-id", appID)
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("response from invokeByAppID method: ", string(body))
}

func invokeByInvokeApiURL() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v1.0/invoke/%s/method/echo?msg=hello-world", address, appID), nil)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("response from invokeByInvokeApiURL method: ", string(body))
}

func invokeByDaprSDK() {
	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.InvokeMethod(context.TODO(), appID, "echo?msg=hello-world", "get")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("response from invokeByDaprSDK method: ", string(res))
}
