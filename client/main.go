package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Message represents a request message sent from the client to the server
type Message struct {
	RequestType string
	Data        string
}

func main() {
	// Parse command line arguments
	var queueService = flag.String("queue-service", "", "The external queue service to use (e.g. rabbitmq)")
	var queueName = flag.String("queue-name", "", "The name of the queue to send the request to")
	var requestType = flag.String("request-type", "", "The type of request to send (e.g. add, remove, get, get-all)")
	var data = flag.String("data", "", "The data to include in the request (if applicable)")

	flag.Parse()

	// Read data from command line
	message := *data

	// Send request to queue
	if *queueService == "rabbitmq" {
		sendRequestToRabbitMQ(*queueName, *requestType, message)
	} else {
		log.Println("Invalid queue service")
	}
}

func sendRequestToRabbitMQ(queueName, requestType, message string) {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		queueName, // name
		true,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Send message to the queue
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("%s:%s", requestType, message)),
		})
	if err != nil {
		log.Fatal("Failed to publish a message:", err)
	}

	log.Printf(" [x] Sent %s:%s", requestType, message)
}