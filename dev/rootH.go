package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"bufio"
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
	
	var	cfgFile = "msgsrvr.cfg"
	HOST, PORT, errFile := getConfig(cfgFile)
	if( errFile != nil ){
		fmt.Println(errFile.Error() )
		fmt.Println("Unable to read configuration from file :"+cfgFile )
		return
		}

	service := "ws://"+HOST+":"+PORT

	if len(os.Args) > 1 {
		service += "/" + os.Args[1]
	} else {
		service += "/"
		}

	//Connect
	conn, err := websocket.Dial(service, "", "http://"+HOST)
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

