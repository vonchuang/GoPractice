package main

import (
	"bufio"
	"fmt"
	"net"
)

type client chan string

var num int = 0
var mssg = make(chan string)
var enter = make(chan client)
var leaver = make(chan client)

//var leaver = make(chan string)

func main() {

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		num++
		fmt.Printf("Client number: %d\n", num)

		go handler(conn)

	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for { // DON'T FORGET for !!!
		select {
		case cli := <-enter:
			clients[cli] = true
		case m := <-mssg:
			for cli := range clients {
				cli <- m
			}
		case cli := <-leaver:
			delete(clients, cli)
		}
	}

}

func clientWriter(conn net.Conn, message <-chan string) {
	for m := range message {
		fmt.Fprintf(conn, "--> %s\n", m)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	var namecheck int = 0
	var username string
	var cli = make(chan string)
	go clientWriter(conn, cli)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		if namecheck == 0 {
			namecheck++
			username = input.Text()
			cli <- "Welcome " + username + "!"      // back to client itself
			mssg <- username + " into the chatrom~" // broadcast mss
			enter <- cli                            // save client channel

			fmt.Printf("User %s into the chatrom\n", username)

		} else {
			mssg <- username + ": " + input.Text()
			fmt.Printf("%s: %s\n", username, input.Text())
		}
	}

	num--
	mssg <- username + " leave the chatroom"
	leaver <- cli
	fmt.Printf("User %s leave the chatrom\n", username)
	fmt.Printf("Client number: %d\n", num)
}
