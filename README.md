# bloxRoute Project
## Introduction

This project is a client-server application written in Golang that satisfies the following requirements:
* multiple-threaded server(channels);
* multiple clients;
* External queue between the clients and server;
* Implement a server with a data structure that holds the data in memory and maintains the order of items as they are added;
* Implement the following client requests: AddItem, RemoveItem, GetItem, GetAllItems;

## Project decisions
I chose to use RabbitMQ as the external queue because it is a good and fast solution for local development, even though it may not be the most efficient or scalable.

I decided to use a simple array of strings to store the data in the server, although I could have used a hash map.

I allowed the client to use commands to interact with the server. 

## to take in consideration
I am not an expert in Golang, and have only used it in online courses and playing with Bor on the Polygon EVM.

This is my first time using RabbitMQ, so I kept the configuration basic.

I have not implemented any tests, as it was not requested, but I could have and would prefer to.

## How to run the project

### Requirements

* Golang
* RabbitMQ
* RabbitMq default configuration(guest:guest)

### First steps
* `git clone https://github.com/ethanolle/bloxRouteClientServer.git`
* `cd bloxRouteClientServer`
* open two terminals in the project folder and run the following commands in each terminal.

### Run the server
* `cd server`
* `go run ./main.go`

### Run the client
* `cd client`
* `go build ./main.go`
* `./main --queue-service=SERVICENAME --queue-name=QUEUENAME --request-type=METHOD(add,remove,getAll,get) --data=STRING`

#### default queue-service and queue-name:
* queue-service: rabbitmq
* queue-name: request_queue

#### request-type options:
* add
* remove
* getAll
* get

#### Examples: 
* `./main --queue-service=rabbitmq --queue-name=request_queue --request-type=add --data=A`
* `./main --queue-service=rabbitmq --queue-name=request_queue --request-type=getAll`

Note that I allowed the client to modify the queue name and service, rather than hardcoding them, to allow the client to communicate with multiple queue-services/queue-names.

## Notes
This project was a fun learning experience for me in using Golang and RabbitMQ. I may not have been able to complete it without the help of online resources, such as chatGpt for debugging/ideas and Stack Overflow for understanding RabbitMQ and installation issues with Homebrew🙃.






