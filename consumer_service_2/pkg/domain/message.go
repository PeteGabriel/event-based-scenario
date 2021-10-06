package domain

import (
	json2 "encoding/json"
	"log"
)

type Message struct {
	ContentType string `json:"content-type"`
	Id string `json:"id"`
	ConsumerTag string `json:"consumer-tag"`
	Content []byte `json:"content"`
}

func (m Message) String() string{
	json, err := json2.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	return string(json)
}