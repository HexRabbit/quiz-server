package fvt

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ccns/quiz-server/config"
)

// VerifyGetProvokes verify the GET method of the route /provokes.
func VerifyGetProvokes() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyGetProvokes")

	fmt.Printf(config.Config.FVT.Section, "All Provokes")
	url := "http://0.0.0.0:8080/v1/provokes"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ GET %s\n```\n", url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))

	fmt.Printf(config.Config.FVT.Section, "Query Provokes by Correctness")
	url = "http://0.0.0.0:8080/v1/provokes?correct=true"
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ GET %s\n```\n", url)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))
}

// VerifyPostProvokes verify the POST method of the route /provokes.
func VerifyPostProvokes() {

	fmt.Printf(config.Config.FVT.Topic, "VerifyPostProvokes")
	url := "http://0.0.0.0:8080/v1/provokes"
	payload := `{
	"correct":true,
	"message":"%s"
}`
	jsonStr := fmt.Sprintf(payload, testProvokeMessage)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(config.Config.FVT.Section, "New Provoke")
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ POST %s\n```\n", url)
	fmt.Printf(config.Config.FVT.Detail, "example payload")
	fmt.Printf("```\n%s\n```\n", jsonStr)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))

	fmt.Printf(config.Config.FVT.Section, "Duplicate Provoke")
	fmt.Printf(config.Config.FVT.Detail, "method and url")
	fmt.Printf("```\n$ POST %s\n```\n", url)
	fmt.Printf(config.Config.FVT.Detail, "example payload")
	fmt.Printf("```\n%s\n```\n", jsonStr)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(config.Config.FVT.Detail, "example response")
	fmt.Printf("```\n%s\n```\n", string(bodyBytes))
}
