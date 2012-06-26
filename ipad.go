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
	
	service := "ws://localhost:9030/read"
	//service := "ws://10.0.1.6:9030/"
	//service := "ws://dev.itelehome.com:9030/"

	//Connect
	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)
	
	//Send	
	smsg := locationUpdate()
	err = websocket.Message.Send(conn, smsg)
	checkError(err)
	
	//Receive
	var rmsg string	
	err = websocket.Message.Receive(conn, &rmsg)
	checkError(err)
	fmt.Println("Received from server: " + rmsg)
	
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func locationUpdate() string {
	//jstr := "{"
	
	loc := Location{}
	loc.Ag = "locmgr"
	loc.Sb = "updts"
	loc.Fr = "HQ"
	loc.DT = "2012-05-02:12:0000"	
	
	jstr := str2js(loc) 
	
	//jstr += str2js(loc)   
	//jstr += "}"
	
	return jstr
}