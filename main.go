package main

import (
	"./broker"
	"fmt"
	"net"
    "time"
        "encoding/base64"

)

func main() {
	fmt.Println("Connecting...")
	conn, err := net.Dial("tcp", "broker.labs.sapo.pt:3323")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected.")
	in := make(chan string)

	go broker.Publisher(conn, "/test", in)
	fmt.Println("Started goroutine.")
	in <- base64.StdEncoding.EncodeToString([]byte("Hello World"))
	fmt.Println("Sent.")
    time.Sleep(100 * time.Millisecond)

	conn.Close()

	/*
			fmt.Printf("Response: %+v\n", res)
		msg := broker.NewPublishMessage("/sapo/test/announcements", &payload)
		payload := "Hello world!"
		b, _ := json.Marshal(*msg)
		fmt.Println(string(b))
		msg = broker.NewSubscribeMessage("/sapo/.*")
		b, _ = json.Marshal(*msg)
		fmt.Println(string(b))
	*/
}
