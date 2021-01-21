package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
	"time"
)

type API int

var promise [5]bool
var ids [5]int

var client1, client2, client3, client4 *rpc.Client

// var err1 *net.OpError
// var err2 *net.OpError
// var err3 *net.OpError
// var err4 *net.OpError

var reply1 string
var reply2 string
var reply3 string
var reply4 string

var finalreply string

var id = 1

func main() {
	ids[4] = id

	fmt.Println("Server5")

	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8085")

	if err != nil {
		log.Fatal("Listener error", err)
	}

	go dailClients()

	log.Printf("serving rpc on port %d", 8085)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

}

func dailClients() {

	fmt.Println("inside dailclients but not slept")
	time.Sleep(time.Second * 10)
	fmt.Println("inside dailclients after sleep")

	client1, _ = rpc.DialHTTP("tcp", "192.168.0.101:8081")
	client2, _ = rpc.DialHTTP("tcp", "192.168.0.101:8082")
	client3, _ = rpc.DialHTTP("tcp", "192.168.0.101:8083")
	client4, _ = rpc.DialHTTP("tcp", "192.168.0.101:8084")
	fmt.Println("after dails calls")

	// if err1 != nil {
	// 	log.Fatal("Connection error: ", err1)
	// }
	// if err2 != nil {
	// 	log.Fatal("Connection error: ", err2)
	// }
	// if err3 != nil {
	// 	log.Fatal("Connection error: ", err3)
	// }
	// if err4 != nil {
	// 	log.Fatal("Connection error: ", err4)
	// }

}

//will send prepare messaages to other servers and if got majority promises then send accept messages to all servers
func (a *API) Proposer(empty string, reply *string) error {

	//if client calls proposer
	/*
		call prepare at server1
		call prepare at server2
		call prepare at server3
		call prepare at server4

	*/
	majortiy := 1

	var sendId = fmt.Sprintf("%d-%d", id, 4)

	client1.Call("API.Prepare", sendId, &reply1)
	if reply1 == "yes" {
		majortiy = majortiy + 1
	}

	client2.Call("API.Prepare", sendId, &reply2)
	if reply2 == "yes" {
		majortiy = majortiy + 1
	}

	client3.Call("API.Prepare", sendId, &reply3)
	if reply3 == "yes" {
		majortiy = majortiy + 1
	}

	client4.Call("API.Prepare", sendId, &reply4)
	if reply4 == "yes" {
		majortiy = majortiy + 1
	}

	fmt.Println("Number of Acceptors who responded yes", majortiy)

	if majortiy >= 3 {
		/*
			If majority promises, send accept command messages to servers 1,2,3,4
		*/
		selfAccept("command5")
		client1.Call("API.Accept", "command5", &reply1)
		client2.Call("API.Accept", "command5", &reply2)
		client3.Call("API.Accept", "command5", &reply3)
		client4.Call("API.Accept", "command5", &reply4)
	}

	id = id + 1
	ids[4] = id
	*reply = finalreply
	time.Sleep(time.Second * 1)
	finalreply = ""
	return nil

}

//when other server calls prepare, will check if already promised or not,
//if not promised anyone then send reply as yes
//if promised someone, check the id of the incoming message
//if the id is less than the current maxid then just send reply as no
//if the id is greater than the current maxid then change the id of promised server to maxid and send reply as no
func (a *API) Prepare(r string, reply *string) error { // r format should be 2-0,2-1,3-2, the 1st number is id,second is to identify the requesting server

	var response string

	req := strings.Split(r, "-")

	id, err := strconv.Atoi(req[0])
	if err != nil {
		fmt.Println("Cannot convert string to int")
	}
	reqServer, err := strconv.Atoi(req[1])
	if err != nil {
		fmt.Println("Cannot convert string to int")
	}
	ids[reqServer] = id

	var maxid int
	maxid = ids[0]
	for i := 0; i < 5; i++ {
		if ids[i] > maxid {
			maxid = ids[i]
		}
	}
	if id >= maxid {
		promised := false
		for i := 0; i < 5; i++ {
			if promise[i] == true {
				promised = true
				break
			}
		}
		//if i didn't promised anyone yet
		//3 0 0 0 0-promise done for p1
		//3 5 0 0 0
		//5 5 0 0 0
		if promised == false {
			promise[reqServer] = true
			response = "yes"
		} else {
			if maxid == ids[reqServer] {
				//to whom i promised that corresponding id should be maxid
				var index int
				for i := 0; i < 5; i++ {
					if promise[i] == true {
						index = i
						break
					}
				}
				ids[index] = maxid
				response = "no"
			} else {
				response = "no"
			}
		}

	} else {
		response = "no"
	}

	*reply = response
	return nil
}

func (a *API) Accept(r string, reply *string) error { // r will contain command
	r += "-4" //appending server number to identify the server
	selfLearn(r)
	client1.Call("API.Learner", r, &reply1)
	client2.Call("API.Learner", r, &reply2)
	client3.Call("API.Learner", r, &reply3)
	client4.Call("API.Learner", r, &reply4)

	*reply = ""

	return nil
}

func selfAccept(r string) {
	r += "-4"
	selfLearn(r)
	client1.Call("API.Learner", r, &reply1)
	client2.Call("API.Learner", r, &reply2)
	client3.Call("API.Learner", r, &reply3)
	client4.Call("API.Learner", r, &reply4)

}

func (a *API) Learner(r string, reply *string) error {
	req := strings.Split(r, "-")
	time.Sleep(time.Second * 2)
	fmt.Println("I got a command ", req[0], " to execute from the server", req[1])
	finalreply = "command " + req[0] + " is executed."
	*reply = ""
	for i := 0; i < 5; i++ {
		promise[i] = false
	}
	return nil
}

func selfLearn(r string) {
	req := strings.Split(r, "-")
	time.Sleep(time.Second * 2)
	fmt.Println("I got a command ", req[0], " to execute from the server", req[1])
	finalreply = "command " + req[0] + " is executed."
	for i := 0; i < 5; i++ {
		promise[i] = false
	}
}
