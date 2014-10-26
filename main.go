package main

import (
	"./broker"
    "encoding/json"
	"fmt"
)

func main() {
	/*
		conn, err := net.Dial("tcp", "broker.labs.sapo.pt:3323")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Response: %+v\n", res)
	*/
	payload := "Hello world!"
	msg := broker.NewPublishMessage("/sapo/test/announcements", &payload)
	b, _ := json.Marshal(*msg)
	fmt.Println(string(b))
	msg = broker.NewSubscribeMessage("/sapo/.*")
	b, _ = json.Marshal(*msg)
	fmt.Println(string(b))
}
