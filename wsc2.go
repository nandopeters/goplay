

/* EchoClient
 */
package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"os"
//	"bytes"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "ws://host:port")
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := websocket.Dial(service, "", "http://localhost")
	checkError(err)
	var msg string 
	var input []byte = make( []byte, 100)	
	for {
		err := websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				// graceful shutdown by server
				break
			}
			fmt.Println("Couldn't receive msg " + err.Error())
			break
		}
		fmt.Println("Received from server: " + msg)
		// read from stdin
		cnt, err := os.Stdin.Read( input)

		inpStr := string (rTrim(input, cnt-1))
		err = websocket.Message.Send(conn, inpStr)

		if err != nil {
			fmt.Println("Couldn't return msg")
			break
		}

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

