package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)

type nsqhandler struct{}

var stp chan bool = make(chan bool)

func (h *nsqhandler) HandleMessage(m *nsq.Message) error {
	defer func() {
		close(stp)
	}()
	if len(m.Body) == 0 {
		return nil
	}
	fmt.Printf("%s \n", m.Body)
	return nil
}

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "ch1", config)
	if err != nil {
		log.Panic(err)
	}
	consumer.AddHandler(&nsqhandler{})
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Panic(err)
	}
	<-stp
	consumer.Stop()
}
