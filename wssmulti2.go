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

type AppChannelQ struct {
	chanQ map[string]chan string
	}
	
var msgQ	AppChannelQ

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

	msgQ.insertMsgAllQ( reply )
	//writeChann(reply)

}

func iPad(ws *websocket.Conn) {
	
	//Receive Message	
	
	var who string;

	err := websocket.Message.Receive(ws, &who)
	checkError(err)
	fmt.Println("Connected to client :", who)
	
	msgQ.addQ( who )
	//msgQ.popQ()
	fmt.Println( "iPad :", msgQ )

	for s := range msgQ.chanQ[who] {
        fmt.Println("Sending to '",who, "': ", s)
        err := websocket.Message.Send(ws, s)
        checkError(err)
        }
     fmt.Println("DONE");
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
//msgQ = make( map[string][]string )
msgQ.initQ()
//msgQ.addQ("DB")
  go msgQ.doDBQ()

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



//----------------------------------

func  (cq *AppChannelQ) initQ ( ) {
	cq.chanQ = make(map[string]chan string)
	}

func  (cq AppChannelQ) addQ ( who string) {
	cq.chanQ[who] =  make(chan string) 
	}

func  (cq *AppChannelQ) popQ ( who string ) string {
	return <- cq.chanQ[who]
}

func  (cq *AppChannelQ) popAllQ ( who string )  {
		for s := range cq.chanQ[who] {
			fmt.Println("pop:",s)
		}
	}
	
func  (cq *AppChannelQ ) insertMsgAllQ( msg string) {
	for k, _ := range cq.chanQ {
		fmt.Println("insertMsgAllQ:  k=",k);
		//cq.chanQ[k] = append( cq.chanQ[k], msg )
		cq.chanQ[k] <- msg
		}
	}
	
func (cq *AppChannelQ) doDBQ() {
	who := "DB"
	cq.addQ(who)
	for s := range msgQ.chanQ[who] {
        fmt.Println("THIS IS DB Q: ", s)
        }
 	}

//----------------------------------------------
func writeChann (msg string) {
	cs <- msg
}

func readChann() {
	for s := range cs {
        fmt.Println("Received msg: ", s)
        }
}

