package consumer

import (
	"database/sql"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
	"queue_assignment/db"
)

type payload struct {
	Counter int `json:"counter"`
	Timestamp int64 `json:"time_stamp"`
}

func Consumer()  {
	log.Println("Consuming")
	channel := ConnectToQueue()
	dbConnection := db.Connecttodb()
	ConsumeFromQueue(channel, dbConnection)
	defer channel.Close()
}

func failOnError(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectToQueue() *amqp.Channel {
	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	log.Println("Connecting to rabbit mq")
	conn, err := amqp.Dial(rabbitmqURL)
	failOnError(err)

	log.Println("Creating a channel")
	ch, err := conn.Channel()
	failOnError(err)

	return ch
}

func ConsumeFromQueue(ch *amqp.Channel, dbConnection *sql.DB) {
	queue := os.Getenv("QUEUE_NAME")
	msgs, err := ch.Consume(queue, "consumer-1", false, false, false, false, nil)
	failOnError(err)
	var data payload
	consumerChan := make(chan bool)
	go func() {
		for msg := range msgs {
			msgBytes := msg.Body
			log.Println("Received ", string(msgBytes))
			if err = json.Unmarshal(msgBytes, &data); err != nil {
				failOnError(err)
			}
			go db.SaveToDb(dbConnection, data.Counter, data.Timestamp)
			_ = msg.Ack(false)
		}
	}()
	<- consumerChan
	log.Println("done")

}
