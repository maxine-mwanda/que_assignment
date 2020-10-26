package main

import (
	"github.com/joho/godotenv"
	"queue_assignment/consumer"
	"queue_assignment/producer"
)

func main () {
	_= godotenv.Load()
	qChan := make(chan bool)
	go consumer.Consumer()
	go producer.Producer()
	<- qChan
}