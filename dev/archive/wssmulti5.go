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


/*
messagetype:
location			payload { schedule_id, session_id, datetime, latitude, longitude, msg }
broadcast			payload { schedule_id, session_id, datetime, msg }
join_session		payload { schedule_id, session_id, datetime,	msg }
leave_session		payload { schedule_id, session_id, datetime,	msg }
session_ended		payload { schedule_id, session_id, datetime, msg }
session_started		payload { schedule_id, session_id, datetime, msg }
session_paused		payload { schedule_id, session_id, datetime, msg }
session_resumed		payload { schedule_id, session_id, datetime, elapsedtime, msg }
chat_msg			payload { schedule_id, session_id, datetime, msg }
list_participants	payload { schedule_id, session_id, participants [] }
query_participants	payload { schedule_id, session_id }
url					payload {session_id, url }
*/


type parts	struct {
	Users	string
	}

type load	struct {
		Schedule_id		string
		Session_id		string
		Latitude		string
		Longitude		string
		Datetime		string
		Elapsedtime		string
		Msg				string
		Participants	[]parts
		Url				string
		}

type Message struct {
Messagetype		string
Broadcast		string	// Y
Key				string
From			string
To				string
Payload			load
}



type AppChannelQ struct {
	chanQ map[string]chan string
	}
	
var msgQ	AppChannelQ

var cs = make(chan string)
var inMsgNo = 1
var outMsgNo = 1


// root service handler.  For testing reflects the message back
func rootH(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	err := websocket.Message.Receive(ws, &reply)
	checkError(err)
	fmt.Println("root msg received:", inMsgNo)
	fmt.Println( reply )
	
	//Send Message
	err = websocket.Message.Send(ws, reply)
	checkError(err)
}




func itelPublishOnce(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	
		err := websocket.Message.Receive(ws, &reply)

		checkError(err)
		fmt.Println("Received Message No:", inMsgNo)
		fmt.Println( reply )
		inMsgNo++
		
		var	m Message
		err1 := json.Unmarshal([]byte(reply), &m)
		checkError(err1)

		err = websocket.Message.Send(ws, "ACK")
		checkError(err)
	
		msgQ.insertMsgAllQ( reply )
	
}

func Publish(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	var	selfNo int
	selfNo = 1
	fmt.Println("ENTERING Publish()");
		for {
			err := websocket.Message.Receive(ws, &reply)
			if (err != nil ) {
				fmt.Println("Exiting Publis.  Error = ", err.Error() );
				return;
				}

		//	fmt.Println("Received Message No:", inMsgNo)
		//	fmt.Println(reply);
			inMsgNo++
			
			fmt.Println("self No:", selfNo)
			selfNo++
			
			var	m Message
			err1 := json.Unmarshal([]byte(reply), &m)
			checkError2(err1)
			fmt.Println("Unmarshalled:", m);
			
		//	err = websocket.Message.Send(ws, "ACK")
		//	checkError(err)
		
			msgQ.insertMsgAllQ( reply )
		}
	fmt.Println("\n Exiting Publish() \n\n");
}


func Subscribe(ws *websocket.Conn) {
	
	//Receive Message	
	// iPad identifies itself by sending it's id (who)
	var who string;

	err := websocket.Message.Receive(ws, &who)
	checkError(err)
	fmt.Println("Connected to client :", who)
	
	// add the connection to the messageQ
	msgQ.addQ( who )
	
	fmt.Println( "iPad :", msgQ )

	for s := range msgQ.chanQ[who] {
        fmt.Println("Sending to '",who, "': ", s)
        err := websocket.Message.Send(ws, s)
        if (err != nil ) {
			fmt.Println("Existing Subscribe.  Error = ", err.Error() );
			return;
			}
        //checkError(err)
        }
     fmt.Println("DONE.  Existing Subscribe");
}





func exit_handler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("exit handler called")
	return
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
	
	_, err = f.WriteString(line)

	f.Close()
}


func main() {

	msgQ.initQ()
	
	go msgQ.doDBQ()

	http.Handle("/Subscribe", websocket.Handler(Subscribe))
	http.Handle("/itelPublishOnce", websocket.Handler(itelPublishOnce))
	http.Handle("/Publish", websocket.Handler(Publish))
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

func checkError2(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())

			}
}
