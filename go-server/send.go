package main

import (
	"log"
	"fmt"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func send(disk DiskStatus) {

    conn, err := amqp.Dial("amqp://ashraf:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	all := fmt.Sprintf("\tAll: %.2f GB\n", float64(disk.All)/float64(GB))
	used := fmt.Sprintf("\tUsed: %.2f GB\n", float64(disk.Used)/float64(GB))
	free:= fmt.Sprintf("\tFree: %.2f GB\n", float64(disk.Free)/float64(GB))

	body := fmt.Sprintf("\n%s%s%s", all, used, free)

	err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
			})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}