package main

import (
	"fmt"
	"encoding/json"
	"os"
	"code.google.com/p/go.net/websocket"
)

type Location struct {
	Ag	string
	Sb	string
	Fr	string
	DT	string
}

func str2js(c Location) string {
	cjs, err := json.Marshal(c)	
	if err != nil {
        return "Json failed"
    }
    return string(cjs)
}

func main() {
	
	var who string
	
	if len(os.Args) > 1 {
		who = os.Args[1]
	} else {
		who = "ALFONSO"
		}
	
	service := "ws://dev.itelehome.com:9030/iPad"


	//Connect
	conn, err := websocket.Dial(service, "", "http://dev.itelehome.com")
	checkError(err)
	
	fmt.Println("Connected");
	
	err = websocket.Message.Send(conn, who)
    checkError(err)
		
	for {
		//Receive
		var rmsg string	
		err = websocket.Message.Receive(conn, &rmsg)
		checkError(err)
		fmt.Println("Received:")
		fmt.Println( rmsg )
		}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

