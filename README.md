# Chatrm 

A simple TCP chat program using a client/server model.

## Description

Chatrm allows various users to communcate with each other over a TCP connection that is facilitated by a server. Each user simply connects to the server,
indicating which individual they would like to enter a chat with. The server maintains a mapping of usernames to net.Conn connections, allowing accurate message routing to the proper destination. The program follows a hub-and-spoke model, where the server maintains one TCP connection with each user. 

## Getting Started

### Dependencies

* This program was developed and tested on Mac OS using go version 1.15.7 darwin/amd64

### Server setup

* Navigate to the program's directory.
* Execute the following commands.
```
go run server.go 
```



## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```
