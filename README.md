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
[enter the port you would like to host chatrm on]
```

![img](/images/server2.png)

### Client setup

* Navigate to the program's directory
* Ensure that a server is running (if testing locally), see previous section
* Execute the following commands
```
go run client.go
```

![img](/images/client1.png)
```
[enter your server IP address]
```
![img](/images/client2.png)
```
[enter your server port number]
```
![img](/images/client3.png)
```
[enter your username]
```
![img](/images/client4.png)
```
[enter the username of the individual you would like to chat with]
```
![img](/images/client6.png)
```
[type your messages and press enter!]
```
![img](/images/client7.png)
![img](/images/client8.png)

# Design Structure

Message format: DESTINATION_SOURCE_MESSAGE

## Server
* Obtain port via user input
* Open a goroutine that will continually listen on that port.
* Accept all incoming connections and associate it with the respective username as provided by the client program. This data is stored in a global map serverside.
* Open a new goroutine for each user.
* Within each of these goroutines, continually accept messages and parse their format for the destination, source, and message content. 
* Call the sendMessage() function which sends that message over the proper channel based on the global username:net.Conn mapping.

## Client
* Obtain the chatrm server IP and port from the user
* Dial into the open port and save the connection
* Spawn a goroutine that will listen on the connection for incoming messages and print them to the screen
* Quietly advertise the client's username to the server in an initial message. 
* Continually read strings from the user, handling any input errors and sending the message to the server with format DESTINATION_SOURCE_MESSAGE

## Help

Any advise for common problems or issues.
```
If the client setup is not working, ensure that the server is properly configured.
If any devices involved in the connection fail, then the server will exit. 
```
