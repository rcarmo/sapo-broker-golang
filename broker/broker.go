package broker

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
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
				ActionType:      "PUBLISH",
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
