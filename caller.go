package main

import (
	"crypto/rand"
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
	fmt.Printf("caller starting.\n")
	replyTo := getUniqReplyAddress()
	produceRequest(replyTo)
	resp := consumeReply(replyTo)
	fmt.Printf("caller done with response='%#v'.\n", resp)
}

func getUniqReplyAddress() string {
	var a [16]byte
	_, err := rand.Read(a[:16])
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("caller-call-number-%x#ephemeral", a[:])
}

func produceRequest(replyTo string) {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	mm := MyMsg{
		Hello:   "hello at " + fmt.Sprintf("%v", time.Now()),
		ReplyTo: replyTo,
	}

	jsonBytes, err := json.Marshal(mm)
	if err != nil {
		panic(err)
	}

	err = w.Publish("write_test", jsonBytes)
	if err != nil {
		log.Panic("Could not connect")
	}

	fmt.Printf("published '%#v'\n", mm)
	w.Stop()
}

func consumeReply(addrRepliedTo string) *MyMsg {

	ch := make(chan *MyMsg)

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(addrRepliedTo, "ch#ephemeral", config)
	if err != nil {
		panic(err)
	}

	q.ChangeMaxInFlight(1)

	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		//log.Printf("Got a message: %#v", message)

		var mm MyMsg
		err := json.Unmarshal(message.Body, &mm)
		if err != nil {
			panic(err)
		}

		log.Printf("caller: Got a message Body (as MyMsg): %#v", mm)

		//q.ChangeMaxInFlight(0)
		q.Stop()
		message.Finish()
		//message.Requeue(1)

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
