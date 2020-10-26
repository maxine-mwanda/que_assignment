package producer

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type payload struct {
	Counter int `json:"counter"`
	Timestamp int64 `json:"time_stamp"`
}

func Producer(){
	var data payload
	log.Println("Starting producer...")
	channel:= ConnectToQueue()

	for i := 0; i <= 1000; i++ {
		data.Counter = i
		data.Timestamp = time.Now().Unix()
		go PublishToExchange(data, channel)
		//time.Sleep(time.Second)
		}

	}


func FailOnError(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectToQueue() *amqp.Channel {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	log.Println("Connecting to rabbitMq")
	conn, err := amqp.Dial(rabbitmqURL)
	FailOnError(err)

	log.Println("Creating a channel")
	ch, err := conn.Channel()
	FailOnError(err)

	return ch
}

func PublishToExchange(data payload, ch *amqp.Channel) (err error) {
	jsonBytes, _ := json.Marshal(data)
	exchange := os.Getenv("QUEUE_EXCHANGE")
	routingKey := os.Getenv("ROUTING_KEY")

	err =ch.Publish (exchange, routingKey, false, false, amqp.Publishing{
		ContentType:     "application/json",
		Body:            jsonBytes,
	})
	FailOnError(err)
	log.Println("Published to queue")
	return
}
