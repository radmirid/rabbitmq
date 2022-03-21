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

	msg := "message"
	if err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		}); err != nil {
		log.Fatal("error: Failed to declare a queue")
	}
}
