package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"encoding/json"
	"code.google.com/p/go.net/websocket"
	"strconv"
)

type Message struct {
		Ag	string
		Sb	string
		Fr	string
		DT	string
	}


/*
messagetype:
location			payload { session_id, datetime, latitude, longitude, msg }
broadcast			payload { session_id, datetime, msg }
join_session		payload { session_id, datetime,	msg }
leave_session		payload { session_id, datetime,	msg }
session_ended		payload { session_id, datetime, msg }
session_started		payload { session_id, datetime, msg }
session_paused		payload { session_id, datetime, msg }
session_resumed		payload { session_id, datetime, elapsedtime, msg }
chat_msg			payload { session_id, datetime, msg }
list_participants	payload { session_id, participants [] }
query_participants	payload { session_id }
url					payload {session_id, url }
*/


type participants	struct {
	Particiapants	string
	}

type Load	struct {
		Session_id		string
		Latitude		string
		Longitude		string
		Datetime		string
		Msg				string
		Participants	participants
		Url				string
		}
				
type MM struct {
	Msgtype		string
	Broadcast	string	// Y
	Key			string
	From		string
	To			string
	Payload		Load
	}




type AppChannelQ struct {
	chanQ map[string]chan string
	}
	
var msgQ	AppChannelQ

var cs = make(chan string)
var inMsgNo = 1
var outMsgNo = 1

func rootH(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	err := websocket.Message.Receive(ws, &reply)
	checkError(err)
	fmt.Println("root msg received:", inMsgNo)
	fmt.Println( reply )

	
	//Send Message
	msg := "ACK"
	err = websocket.Message.Send(ws, msg)
	checkError(err)
	
}



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
}

func itelMessages(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	//for{
		err := websocket.Message.Receive(ws, &reply)
		checkError(err)
		fmt.Println("Received Message No:", inMsgNo)
		fmt.Println( reply )
		inMsgNo++
		
		var	m MM
		err1 := json.Unmarshal([]byte(reply), &m)
		checkError(err1)
		fmt.Println("Unmarshalled m:", m);
		if m.Msgtype=="join_session" {
			fmt.Println("Msgtype:",m.Msgtype);
			}
		msg := "ACK"

		err = websocket.Message.Send(ws, msg)
		checkError(err)
	
		msgQ.insertMsgAllQ( reply )
	//}
}

func itelMessages0(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	who := "DB";

	err := websocket.Message.Receive(ws, &reply)
	checkError(err)
	fmt.Println("Connected to client :", reply)
	msgQ.insertMsgAllQ( reply )
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


func iPad2(ws *websocket.Conn) {
	
	//Receive Message	
	
	who := "DB";
/*
	err := websocket.Message.Receive(ws, &who)
	checkError(err)
	fmt.Println("Connected to client :", who)
*/	
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
	http.Handle("/iPad2", websocket.Handler(iPad2))
	http.Handle("/itelMessages", websocket.Handler(itelMessages))
	http.HandleFunc("/exit", exit_handler)
	http.Handle("/", websocket.Handler(rootH))
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
	fName := "DB.log"
	for s := range msgQ.chanQ[who] {
		line := "DB Msg: " + strconv.Itoa(inMsgNo) + s +"\n"
		myFileAppendLine(fName, line )
        }
 	}


//-------------------
func myFileAppendLine ( fName string, line string )  {
	
	f, err := os.OpenFile(fName, os.O_WRONLY , 644)
	if( err != nil ) {
		f, err = os.Create(fName)
		checkError(err)
		}else {
			_, err = f.Seek(0, os.SEEK_END)
		}
	fmt.Println("*File:", f )
	
	_, err = f.WriteString(line)

	f.Close()
}
