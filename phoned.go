package main

import (
	"fmt"
	"encoding/json"
	"os"
	"code.google.com/p/go.net/websocket"
	"time"
)

type Location struct {
	Ag		string
	Sb		string	
	Fr		string
	DT		string
	Lat		string
	Lng		string
	Msg		string
}

func str2js(c Location) string {
	cjs, err := json.Marshal(c)	
	if err != nil {
        return "Json failed"
    }
    return string(cjs)
}

func main() {
	
	service := "ws://localhost:9030"

	if len(os.Args) > 1 {
		service += "/" + os.Args[1]
	} else {
		service += "/phone"
		}

	//Connect
	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)


	//Send
	smsg := locationUpdate( )
	err = websocket.Message.Send(conn, smsg)
	checkError(err)

	//Receive
	var rmsg string	
	err = websocket.Message.Receive(conn, &rmsg)
	checkError(err)
	fmt.Println("Received: " + rmsg)
	
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
	loc.Sb = "locupdt"
	loc.Fr = "4"
	loc.DT = time.Now().String()
	loc.Lat = "60.00"
	loc.Lng = "-120.00"
	loc.Msg = "Arrived"
	
	jstr := str2js(loc)   
	//jstr += "}"
	
	return jstr
}