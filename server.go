package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var chats = make(map[string]net.Conn)
var outbound = make(map[string]string)

// Send the message over the proper connection based on the username : connection global map.
func sendMessage(to string, from string, message string, doneSending chan string) {
	msg := to + "_" + from + "_" + message
	tmpConn := chats[to]
	tmpConn.Write([]byte(msg))
}

// Handles the communication with each client.
// Messages are parsed for their DESTINATION_SOURCE_MESSAGE content, which is then passed to sendMessage()
// An initializing message is always received first by the server to assist in populating the map.
func newClient(con net.Conn, user chan string) {
	reader := bufio.NewReader(con)
	data, err := reader.ReadString('\n')
	handle(err)
	init := strings.Split(strings.TrimSpace(string(data)), "_")[0]
	port := strings.Split(strings.TrimSpace(string(data)), "_")[1]
	user <- init
	user <- strings.Split(port, "\n")[0]
	for {
		data, err := reader.ReadString('\n')
		handle(err)
		if strings.Split(string(data), "\n")[0] == "EXIT" {
			return
		}

		msgSections := strings.Split(strings.TrimSpace(string(data)), "_")
		to := msgSections[0]
		from := msgSections[1]
		fullMsg := msgSections[2] + "\n"
		doneSending := make(chan string)
		sendMessage(from, to, fullMsg, doneSending)
	}
}

// Prints ASCII art 'chatrm' banner
func printBanner() {
	fmt.Println("       .__            __                   __                                           \n  ____ |  |__ _____ _/  |________  _____   \\ \\    ______ ______________  __ ___________ \n_/ ___\\|  |  \\\\__  \\\\   __\\_  __ \\/     \\   \\ \\  /  ___// __ \\_  __ \\  \\/ // __ \\_  __ \\\n\\  \\___|   Y  \\/ __ \\|  |  |  | \\/  Y Y  \\  / /  \\___ \\\\  ___/|  | \\/\\   /\\  ___/|  | \\/\n \\___  >___|  (____  /__|  |__|  |__|_|  / /_/  /____  >\\___  >__|    \\_/  \\___  >__|   \n     \\/     \\/     \\/                  \\/            \\/     \\/                 \\/       ")
}

// Panic on any errors
func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// Continually listen on the chatrm port
// Accept any new connections and handle the client in a separate goroutine
func listen(host string, port string, terminate chan string) {
	l, err := net.Listen("tcp", host)
	handle(err)
	defer l.Close()
	fmt.Printf("Succesfully hosting chatrm on localhost:%s\n", port)
	for {
		user := make(chan string)
		con, err := l.Accept()
		handle(err)
		go newClient(con, user)
		tmpUser := <-user
		port := <-user
		chats[tmpUser] = con
		outbound[tmpUser] = port

		fmt.Println(chats)
	}
	terminate <- "x"
}

// Main driver function.
// Obtain user input and open the listener in a separate goroutine.
func main() {
	terminate := make(chan string)
	printBanner()
	var port string
	ip := "127.0.0.1"

	fmt.Println("Please indicate what port you'd like to host the service on.")
	fmt.Print("chatrm>server>> port ")
	fmt.Scan(&port)
	host := fmt.Sprintf("%s:%s", ip, port)
	go listen(host, port, terminate)
	<-terminate
}
