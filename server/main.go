package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/streadway/amqp"
)

// Message represents a message to be sent or received from the queue
type Message string

type DataStructure interface {
	Add(item string) error
	Remove(item string) error
	Get(item string) (string, error)
	GetAll() ([]string, error)
}

type concreteDataStructure struct {
	items []string
}

func (d *concreteDataStructure) Add(item string) error {
	// Add the item to the end of the slice
	d.items = append(d.items, item)
	return nil
}

func (d *concreteDataStructure) Remove(item string) error {
	// Find the index of the item in the slice
	index := -1
	for i, v := range d.items {
		if v == item {
			index = i
		}
	}

	// Return an error if the item was not found
	if index == -1 {
		return fmt.Errorf("item not found")
	}

	// Remove the item from the slice
	d.items = append(d.items[:index], d.items[index+1:]...)
	return nil
}

func (d *concreteDataStructure) Get(item string) (string, error) {
	// Find the item in the slice
	for _, v := range d.items {
		if v == item {
			return v, nil
		}
	}

	// Return an error if the item was not found
	return "", fmt.Errorf("item not found")
}

func (d *concreteDataStructure) GetAll() ([]string, error) {
	// Return a copy of the slice
	return append([]string{}, d.items...), nil
}

var dataStructure DataStructure = &concreteDataStructure{}

func processRequest(requestMessage Message) (Message, error) {
	// Split the request message into request type and data
	parts := strings.SplitN(string(requestMessage), ":", 2)
	requestType := parts[0]
	data := parts[1]

	// Process the request
	switch requestType {
	case "add":
		// Add the item to the data structure
		err := dataStructure.Add(data)
		if err != nil {
			return "", err
		}
		return "item added", nil
	case "remove":
		// Remove the item from the data structure
		err := dataStructure.
		Remove(data)
		if err != nil {
			return "", err
		}
		return "item removed", nil
	case "get":
		// Get the item from the data structure
		item, err := dataStructure.Get(data)
		if err != nil {
			return "", err
		}
		return Message(item), nil
	case "getAll":
		// Get all items from the data structure
		items, err := dataStructure.GetAll()
		if err != nil {
			return "", err
		}
		return Message(strings.Join(items, ",")), nil
	default:
		return "", fmt.Errorf("invalid request type")
	}
}

func main() {
	// Connect to the queue service
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to queue service: ", err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel: ", err)
	}
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		"request_queue", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare a queue: ", err)
	}

	// Set up a consumer
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
		log.Fatal("Failed to register a consumer: ", err)
	}

	// Server loop
	for msg := range msgs {
		go func(msg amqp.Delivery) {
			// Parse the message
			requestMessage := Message(msg.Body)

			// Process the request
			responseMessage, err := processRequest(requestMessage)
			if err != nil {
				log.Println("Failed to process request: ", err)
				return
			}

			// Write the response to the log file
			fmt.Println(responseMessage)
		}(msg)
	}
}
