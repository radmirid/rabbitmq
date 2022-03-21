package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("error: Failed to connect to RabbitMQ")
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("error: Failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"content", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal("error: Failed to declare a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal("error: Failed to register a consumer")
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received Message: %s", d.Body)
		}
	}()

	log.Printf(" | Waiting for Messages...")

	<-forever
}
