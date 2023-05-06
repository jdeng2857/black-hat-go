package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Status struct {
	Message string
	Status  string
}

func main() {
	resp, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	resp.Body.Close()

	res, err := http.Post(
		"http://IP:PORT/ping",
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	log.Printf("%s -> %s\n", status.Status, status.Message)
}
