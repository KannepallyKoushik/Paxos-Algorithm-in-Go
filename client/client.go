package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

var client1, err1 = rpc.DialHTTP("tcp", "192.168.0.101:8081")
var client2, err2 = rpc.DialHTTP("tcp", "192.168.0.101:8082")
var client3, err3 = rpc.DialHTTP("tcp", "192.168.0.101:8083")
var client4, err4 = rpc.DialHTTP("tcp", "192.168.0.101:8084")
var client5, err5 = rpc.DialHTTP("tcp", "192.168.0.101:8085")

var reply1 string
var reply2 string
var reply3 string
var reply4 string
var reply5 string

func main() {

	if err1 != nil {
		log.Fatal("Connection error: ", err1)
	}
	if err2 != nil {
		log.Fatal("Connection error: ", err2)
	}
	if err3 != nil {
		log.Fatal("Connection error: ", err3)
	}
	if err4 != nil {
		log.Fatal("Connection error: ", err4)
	}
	if err5 != nil {
		log.Fatal("Connection error: ", err5)
	}

	fmt.Println("Client")
	for true {
		fmt.Println("\nOptions:")
		fmt.Println("1.Send a command  \n 2.Exit")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter a server to which you want to send command(total 5 servers)")
			var server int
			fmt.Scanln(&server)

			switch server {
			case 1:
				client1.Call("API.Proposer", "", &reply1)
				fmt.Println(reply1)
			case 2:
				client2.Call("API.Proposer", "", &reply2)
				fmt.Println(reply2)
			case 3:
				client3.Call("API.Proposer", "", &reply3)
				fmt.Println(reply3)
			case 4:
				client4.Call("API.Proposer", "", &reply4)
				fmt.Println(reply4)
			case 5:
				client5.Call("API.Proposer", "", &reply5)
				fmt.Println(reply5)
			}

		case 2:
			fmt.Println()
			fmt.Println("Thank you for using our Application")
			os.Exit(3)
		}
	}
}
