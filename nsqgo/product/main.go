package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

func main() {
	config := nsq.NewConfig()
	producter, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatalln(err)
	}
	msg := []byte("hallo 1")
	topic := "test"
	err = producter.Publish(topic, msg)
	if err != nil {
		log.Fatal(err)
	}
	producter.Stop()
}
