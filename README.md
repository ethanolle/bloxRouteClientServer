# bloxRoute Project
## Introduction

As part of the interview process to bloxRoute i was asked to implement a client-server application using Golang. The project was to implement a client-server application with the following requirements:

* multiple-threaded server(channels);
* multiple clients;
* External queue between the clients and server;
* Implement a server with a data structure that holds the data in memory and maintains the order of items as they are added;
* Implement the following client requests: AddItem, RemoveItem, GetItem, GetAllItems;

## Project decisions
I descided to use RabbitMQ as the external queue because it is a good/fast solution for local development and not because of efficiency or scalability.

I descided to use a simply array of string to make this project. I could have use an hash map but i wanted to keep it simple.

I descided to enable the client to use the command to play with the server I could have done more but again my main goal was to keepItSimple.


## to take in consideration

I'm not a Golang expert I have use it only in online courses and playing with Bor on polygon evm.

It's my first time using rabbitmq so I have kept all the basic configuration.

I have not implemented any test it was not requested but I could have done it and i prefer ;p.

## How to run the project

### Requirements

* Golang
* RabbitMQ

### Run the server
* `cd server`
* `go run ./main.go`

### Run the client
* `cd client`
* `go build ./main.go`
* `./main --queue-service=SERVICENAME --queue-name=QUEUENAME --request-type=METHOD(add,remove,getAll,get) --data=STRING`
Examples: 
`./main --queue-service=rabbitmq --queue-name=request_queue --request-type=add --data=A`
`./main --queue-service=rabbitmq --queue-name=request_queue --request-type=getAll`
Note that I have descided to let the client modify the queue name and the service name. I could have hard coded it but I wanted to let the client be able to talk to multiple servers/queues.





