/* EchoServer
 */
package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"strings"
	// "io"
	"code.google.com/p/go.net/websocket"
)

func Echo(ws *websocket.Conn) {
	fmt.Println("Echoing")
	var reply string
	var msg string
	for n := 0; n < 10; n++ {

			msg = string(n+48) + ": Hello"
	
		    if n > 0 {
				msg =  string(n+48) + ": I tell you what, you said to me, '" + reply + "'"
		}
		//	msg :=  string(n+48) + ": I tell you what, you said to me [ " + reply + " ]"
	
		fmt.Println("Sending to client: " + msg)
		err := websocket.Message.Send(ws, msg )
		if err != nil {
			fmt.Println("Can't send")
			break
		}

		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
		if ( strings.Contains(reply, "exit") ){
			fmt.Println("Leaving now!!")
			os.Exit(0);
		}
	}
}

func socket1(ws *websocket.Conn) {
	var reply string
	fmt.Println("Echoing from  socket1")
	err := websocket.Message.Send(ws,"echo from socket1" )
		if err != nil {
			fmt.Println("Can't send")
		}

	err = websocket.Message.Receive(ws, &reply)
	if err != nil {
		fmt.Println("Can't receive")
	}
	fmt.Println("Received from client: " + reply)
		
	fmt.Println("Sending to client :" + reply)
	msg := "got your message: " + reply
	err = websocket.Message.Send(ws, msg )
	if err != nil {
		fmt.Println("Can't send")
	}
}

func my_http_handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "%s", "you have reached the http handler")
}

func exit_handler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("exit handler called")
	return
}

func main() {

	http.Handle("/", websocket.Handler(Echo))
	http.Handle("/socket1", websocket.Handler(socket1) )
	http.HandleFunc("/http", my_http_handler)
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

