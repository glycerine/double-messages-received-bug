package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

type MyMsg struct {
	Hello     string
	ReplyTo   string
	Reply     string
	ReplyFrom string
}

func main() {
	request := consumeRequest()
	request.Reply = "replying at " + fmt.Sprintf("%v", time.Now())
	request.ReplyFrom = "callee!"
	produceReply(request)
}

func produceReply(mm *MyMsg) {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	jsonBytes, err := json.Marshal(mm)
	if err != nil {
		panic(err)
	}
	err = w.Publish(mm.ReplyTo, jsonBytes)
	if err != nil {
		log.Panic("Could not connect")
	}

	fmt.Printf("callee.ProduceReply(): replied-to: '%s' with value '%#v'\n", mm.ReplyTo, mm)
	w.Stop()
}

func consumeRequest() *MyMsg {

	ch := make(chan *MyMsg)

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer("write_test", "ch", config)
	if err != nil {
		panic(err)
	}

	// if we make Conns exported, no difference. i.e. there
	// is no difference if we comment in/out the following code:
	/*
		conns := q.Conns()
		rdy := 1
		for _, conn := range conns {
			err := conn.WriteCommand(nsq.Ready(rdy))
			if err != nil {
				panic(err)
			}
			rdy = 0 // only the first conn gets RDY=1, the rest get RDY=0
		}
	*/
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		//log.Printf("Got a message: %#v", message)

		var mm MyMsg
		err := json.Unmarshal(message.Body, &mm)
		if err != nil {
			panic(err)
		}

		log.Printf("callee.ConsumeRequest(): Got a message Body (as MyMsg): %#v", mm)
		//log.Printf("sleeping 1 sec")
		//time.Sleep(1 * time.Duration(time.Second))

		message.Finish()

		ch <- &mm
		return nil
	}))

	err = q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}

	mm2 := <-ch

	fmt.Printf("stats = %#v\n", q.Stats())

	return mm2
}
