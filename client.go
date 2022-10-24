package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var listenport string

// Prints the chatrm>client program ASCII art banner
func printBanner() {
	fmt.Println("       .__            __                   __          .__  .__               __   \n  ____ |  |__ _____ _/  |________  _____   \\ \\    ____ |  | |__| ____   _____/  |_ \n_/ ___\\|  |  \\\\__  \\\\   __\\_  __ \\/     \\   \\ \\ _/ ___\\|  | |  |/ __ \\ /    \\   __\\\n\\  \\___|   Y  \\/ __ \\|  |  |  | \\/  Y Y  \\  / / \\  \\___|  |_|  \\  ___/|   |  \\  |  \n \\___  >___|  (____  /__|  |__|  |__|_|  / /_/   \\___  >____/__|\\___  >___|  /__|  \n     \\/     \\/     \\/                  \\/            \\/             \\/     \\/      ")
}

// Panic on all err values
func handle(err error) {
	if err != nil {
		panic(err)
	}
}

//
func newChat(conn net.Conn, username string, destUsername string) {
	rawUsername := username
	username = username + "_" + listenport + "\n"
	// Advertise the client's username to the server so they can associate it with the session
	conn.Write([]byte(username))
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		handle(err)
		msg := rawUsername + "_" + destUsername + "_" + line
		conn.Write([]byte(msg))
	}
}

func listenForMessage(con net.Conn) {
	reader := bufio.NewReader(con)
	for {
		data, err := reader.ReadString('\n')
		handle(err)
		msgSections := strings.Split(strings.TrimSpace(string(data)), "_")
		from := msgSections[1]
		fullMsg := msgSections[2]
		fmt.Printf("[%s>>] %s\n", from, fullMsg)
	}
}

func main() {
	printBanner()
	var server_ip string
	var server_port string
	var username string
	var dest_username string

	fmt.Println("Please enter the chatrm server ip.")
	fmt.Print("chatrm>client>> ")
	fmt.Scan(&server_ip)
	fmt.Println("Please enter the chatrm server port.")
	fmt.Print("chatrm>client>> ")
	fmt.Scan(&server_port)

	fmt.Printf("Connecting to CHATRM>%s:%s\n", server_ip, server_port)
	con, err := net.Dial("tcp", server_ip+":"+server_port)
	handle(err)
	fmt.Println("Connected.")

	fmt.Print("chatrm>client>> Enter your username: ")
	fmt.Scan(&username)
	fmt.Printf("chatrm>%s>> Enter chat session with user: \n", username)
	fmt.Scan(&dest_username)
	go listenForMessage(con)
	newChat(con, username, dest_username)
}
