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



