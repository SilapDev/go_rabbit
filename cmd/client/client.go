package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal("failed to connect to rabbitmq")
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal("failed to open a channel")
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal("failed to declare a queue")
	}

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf(" Golang is the simple: %d ", i)

		if err := ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			}); err != nil {
			log.Fatal("failed to declare a queue")
		}
	}

}
