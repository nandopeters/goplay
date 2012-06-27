

/* EchoClient
 */
package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"os"
	"strings"
//	"bytes"
)

func main() {

	var service string
	
	if len(os.Args) > 1 {
		service = os.Args[1]
	} else {
		service = "ws://localhost:9030"
		}
	
	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)
	var msg string 
	var input []byte = make( []byte, 100)	
	err = websocket.Message.Receive(conn, &msg)
	if err != nil {
		if err == io.EOF {
			// graceful shutdown by server
			os.Exit(0);
		}
		fmt.Println("Couldn't receive msg " + err.Error())
		os.Exit(0);
	}
	fmt.Println("Received from server: " + msg)
	
	for {

		// read from stdin
		cnt, err := os.Stdin.Read( input)
		mystr := string( rTrim(input, cnt-1) )
		if ( strings.Contains(mystr, "exit") ){
			fmt.Println("Leaving now!!")
			os.Exit(0);
		}
		inpStr := string (rTrim(input, cnt-1))
		err = websocket.Message.Send(conn, inpStr)
		if err != nil {
			fmt.Println("Couldn't return msg")
			break
		}
		
		err = websocket.Message.Receive(conn, &msg)
		checkError(err)
		fmt.Println("Received from server: " + msg)
		

	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func rTrim( input []byte, cnt int) ([]byte){
	for i:= cnt; i < len(input) ; i++ {
		input[i] = 0
	} 
	return input
}

