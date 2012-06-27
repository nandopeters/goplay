package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"encoding/json"
	"code.google.com/p/go.net/websocket"
)

type Message struct {
		Ag	string
		Sb	string
		Fr	string
		DT	string
	}

var cs = make(chan string)
var inMsgNo = 1
var outMsgNo = 1

func phone(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	err := websocket.Message.Receive(ws, &reply)
	checkError(err)
	fmt.Println("Received Message No:", inMsgNo)
	fmt.Println( reply )
	inMsgNo++
	
	//Send Message
	msg := "ACK"
	err = websocket.Message.Send(ws, msg)
	checkError(err)
	
	m := decode(reply)
	fmt.Println("Unmarshalled:  ")
	fmt.Println( m, "\n" );

	writeChann(reply)

}



func iPad(ws *websocket.Conn) {
	
	//Receive Message	
	
	var who string;

	err := websocket.Message.Receive(ws, &who)
	checkError(err)
	fmt.Println("Connected to client :", who)

	for s := range cs {
        fmt.Println("Sending: ", s)
        err := websocket.Message.Send(ws, s)
        checkError(err)
        }
	

}



func writeChann (msg string) {
	cs <- msg
}

func readChann() {
	for s := range cs {
        fmt.Println("Received msg: ", s)
        }
}

func decode(jmsg string) Message {
	var m Message
	br := []byte(jmsg)
	err := json.Unmarshal(br, &m)
	if err != nil {
		fmt.Println("Unmarshall error")
	}
	
	return m
}

func locmgr_agt(m Message) {
	switch m.Sb {
		case "locupdt":
			locmgr_agt(m)
		case "updts":
			locmgr_agt(m)
		default:
			fmt.Println("Unknown Subject")
		}
}

func exit_handler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("exit handler called")
	return
}

func main() {


	http.Handle("/phone", websocket.Handler(phone))
	http.Handle("/iPad", websocket.Handler(iPad))
	http.HandleFunc("/exit", exit_handler)
	err := http.ListenAndServe(":9030", nil)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

