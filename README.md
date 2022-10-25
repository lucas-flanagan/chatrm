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


## Help

Any advise for common problems or issues.
```
If the client setup is not working, ensure that the server is properly configured.
Once a chat has been established, it must be maintained. You cannot exit smoothly. 
If any devices involved in the connection fail, then the server will exit. 
```
