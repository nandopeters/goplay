package main

import (
	"fmt"
	"strings"
	"net/http"
	"log"
	"bufio"
	"os"
	"encoding/json"
	"code.google.com/p/go.net/websocket"
	"strconv"
)


/*
messagetype:
location			payload { session_id, datetime, latitude, longitude, msg }
broadcast			payload { session_id, datetime, msg }
join_session		payload { session_id, datetime,	msg }
leave_session		payload { session_id, datetime,	msg }
url					payload {session_id,  url }
app_function		payload { session_id, datetime, msg }
session_ended		payload { session_id, datetime, msg }
session_started		payload { session_id, datetime, msg }
session_paused		payload { session_id, datetime, msg }
session_resumed		payload { session_id, datetime, elapsedtime, msg }
chat_msg			payload { session_id, datetime, msg }
list_participants	payload { session_id, participants [] }
query_participants	payload { session_id }
*/


type parts	struct {
	Users	string
	}

type load	struct {
		Session_id		string
		Schedule_id		string
		Full_name		string
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
	if (err != nil ) {
		fmt.Println("rootH() Receive Error = ", err.Error() );
		return;
		}

	fmt.Println("root msg received:", inMsgNo)
	fmt.Println( reply )
	
	//Send Message
	err = websocket.Message.Send(ws, reply)
	if (err != nil ) {
		fmt.Println("rootH() Send Error = ", err.Error() );
		return;
		}

}




func itelPublishOnce(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	
		err := websocket.Message.Receive(ws, &reply)
		if (err != nil ) {
			fmt.Println("itelPublishOnce() Receive Error = ", err.Error() );
			return;
			}


		fmt.Println("Received Message No:", inMsgNo)
		fmt.Println( reply )
		inMsgNo++
		
		var	m Message
		err1 := json.Unmarshal([]byte(reply), &m)
		checkError2(err1)

		err = websocket.Message.Send(ws, "ACK")
		if (err != nil ) {
			fmt.Println("itelPublishOnce() Send ACK Error = ", err.Error() );
			return;
			}
	
		msgQ.insertMsgAllQ( reply )
	
}


func Publish(ws *websocket.Conn) {
	
	//Receive Message	
	var reply string
	
		for {
			err := websocket.Message.Receive(ws, &reply)
			if (err != nil ) {
				fmt.Println("Publish() Error = ", err.Error() );
				return;
				}

			fmt.Println("Received Message No:", inMsgNo)
			fmt.Println(reply);
			inMsgNo++
			
			var	m Message
			err1 := json.Unmarshal([]byte(reply), &m)
			checkError2(err1)
			fmt.Println("Unmarshalled:", m.To);
			
		//	err = websocket.Message.Send(ws, "ACK")
		//	checkError2(err)
		if ( strings.Contains(m.Messagetype,"broadcast") ){ 
			msgQ.insertMsgAllQ( reply )
			} else {
				msgQ.pushQ(m.To, reply)
			}
		}
}

func PubSubSub(ws *websocket.Conn,  who string) {

	fmt.Println("Inside PubSubSub()");
	
	for s := range msgQ.chanQ[who] {
        fmt.Println("Sending to '",who, "': ", s)
        err := websocket.Message.Send(ws, s)
        if (err != nil ) {
			fmt.Println("Publish() Error = ", err.Error() );
			break;
			}
        }
     fmt.Println("Subsribe() Done for: ", who);
}



func PubSubPub( ws *websocket.Conn) {
	//Receive Message	
	fmt.Println("Inside PubSubPub()");
	var reply string
	
		for {
			err := websocket.Message.Receive(ws, &reply)
			if (err != nil ) {
				fmt.Println("PubSubPub() Error = ", err.Error() );
				return;
				}

			fmt.Println("Received Message No:", inMsgNo)
			fmt.Println(reply);
			inMsgNo++
			
			var	m Message
			err1 := json.Unmarshal([]byte(reply), &m)
			checkError2(err1)
			fmt.Println("Unmarshalled:", m.To);
			
		//	err = websocket.Message.Send(ws, "ACK")
		//	checkError2(err)
		if ( strings.Contains(m.Messagetype,"broadcast") ){ 
			msgQ.insertMsgAllQ( reply )
			} else {
				msgQ.pushQ(m.To, reply)
			}
		}
}

func PubSub( ws *websocket.Conn) {
	//Receive Message	
	// iPad identifies itself by sending it's id (who)
	var who string;
	fmt.Println("starting PubSub()");
	err := websocket.Message.Receive(ws, &who)
	
	if (err != nil ) {
		fmt.Println("Publish() Error receiing WHO on connection ", err.Error() );
		return;
		}
	
	fmt.Println("Connected to client :", who)
	
	// add the connection to the messageQ
	msgQ.addQ( who );
	go PubSubPub( ws );
	go PubSubSub( ws, who );
	// make channel to wait foreve
	c := make(chan int) 
	<- c
	fmt.Println("ENDING PubSub()");
}


func Subscribe(ws *websocket.Conn) {
	
	//Receive Message	
	// iPad identifies itself by sending it's id (who)
	var who string;

	err := websocket.Message.Receive(ws, &who)
	
	if (err != nil ) {
		fmt.Println("Publish() Error receiing WHO on connection ", err.Error() );
		return;
		}
	
	fmt.Println("Connected to client :", who)
	
	// add the connection to the messageQ
	msgQ.addQ( who )
	
	for s := range msgQ.chanQ[who] {
        fmt.Println("Sending to '",who, "': ", s)
        err := websocket.Message.Send(ws, s)
        if (err != nil ) {
			fmt.Println("Publish() Error = ", err.Error() );
			break;
			}
        }
     fmt.Println("Subsribe() Done for: ", who);
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
	
func  (cq *AppChannelQ ) pushQ( who string, msg string) {

		cq.chanQ[who] <- msg

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

	var	cfgFile = "msgsrvr.cfg"
	_, PORT, errFile := getConfig(cfgFile)
	if( errFile != nil ){
		fmt.Println(errFile.Error() )
		fmt.Println("Unable to read configuration from file :"+cfgFile )
		return
		}



	msgQ.initQ()
	
	go msgQ.doDBQ()

	http.Handle("/PubSub", websocket.Handler(PubSub))
	http.Handle("/Subscribe", websocket.Handler(Subscribe))
	http.Handle("/itelPublishOnce", websocket.Handler(itelPublishOnce))
	http.Handle("/Publish", websocket.Handler(Publish))
	http.HandleFunc("/exit", exit_handler)
	http.Handle("/", websocket.Handler(rootH))
	err := http.ListenAndServe(":"+PORT, nil)
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

func getConfig ( cfgFile string) (host string, port string, errOut error )  {
	file, err := os.Open(cfgFile) // For read access.
	if err != nil {
		return "","", err
	}
	
	r := bufio.NewReader(file)
	line, _, err := r.ReadLine()
	aa:= strings.Split(string(line[:]), "=")
	for i := 1; err == nil ; i++ {
		aa = strings.Split(string(line[:]), "=")
		switch {
		case aa[0] == "PORT":
			port=aa[1]
		case aa[0] == "HOST" :
			host = aa[1]		
		}

		line, _, err = r.ReadLine()
		}
	file.Close();

	return host, port, nil
}

