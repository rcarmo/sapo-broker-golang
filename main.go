package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

type (
	BrokerMessagePayload struct {
		Payload *string `json:"payload,omitempty"`
	}
	BrokerActionInfo struct {
		Destination     *string               `json:"destination"`
		ActionType      string                `json:"action_type"`
		DestinationType string                `json:"destination_type"`
		Message         *BrokerMessagePayload `json:"message,omitempty"`
	}
	BrokerAction struct {
		PublishInfo   *BrokerActionInfo `json:"publish,omitempty"`
		SubscribeInfo *BrokerActionInfo `json:"subscribe,omitempty"`
	}
	BrokerMessage struct {
		Action *BrokerAction `json:"action"`
	}
)

func NewPublishMessage(topic string, message *string) *BrokerMessage {
	return &BrokerMessage{
		Action: &BrokerAction{
			PublishInfo: &BrokerActionInfo{
				ActionType:      "SUBSCRIBE",
				Destination:     &topic,
				DestinationType: "TOPIC",
				Message: &BrokerMessagePayload{
					Payload: message,
				},
			},
		},
	}
}

func NewSubscribeMessage(topic string) *BrokerMessage {
	return &BrokerMessage{
		Action: &BrokerAction{
			SubscribeInfo: &BrokerActionInfo{
				ActionType:      "SUBSCRIBE",
				Destination:     &topic,
				DestinationType: "TOPIC",
				Message:         nil,
			},
		},
	}
}

func pack(msg *BrokerMessage) []byte {
	b, _ := json.Marshal(*msg)
	buf := new(bytes.Buffer)
	var data = []interface{}{
		uint16(3),     // JSON transport
		uint16(0),     // version
		int32(len(b)), // payload length
		b,
	}

	for _, item := range data {
		binary.Write(buf, binary.BigEndian, item)
	}
	return buf.Bytes()
}

func main() {
	/*
		conn, err := net.Dial("tcp", "broker.labs.sapo.pt:3323")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Response: %+v\n", res)
	*/
	payload := "Hello world!"
	msg := NewPublishMessage("/sapo/test/announcements", &payload)
	b, _ := json.Marshal(*msg)
	fmt.Println(string(b))
	msg = NewSubscribeMessage("/sapo/.*")
	b, _ = json.Marshal(*msg)
	fmt.Println(string(b))
	b = pack(msg)
	fmt.Printf("%x\n", b)
}
