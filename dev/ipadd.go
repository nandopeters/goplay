package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"bufio"
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
	var	cfgFile = "msgsrvr.cfg"
	HOST, PORT, errFile := getConfig(cfgFile)
	if( errFile != nil ){
		fmt.Println(errFile.Error() )
		fmt.Println("Unable to read configuration from file :"+cfgFile )
		return
		}

	service := "ws://"+HOST+":"+PORT+"/iPad"
	
	var who string
	
	if len(os.Args) > 1 {
		who = os.Args[1]
	} else {
		who = "ALFONSO"
		}
	


	//Connect
	conn, err := websocket.Dial(service, "", "http://"+HOST)
	checkError(err)
	
	fmt.Println("Connected");
	
	err = websocket.Message.Send(conn, who)
    checkError(err)
		
	for {
		//Receive
		var rmsg string	
		err = websocket.Message.Receive(conn, &rmsg)
		checkError(err)
		fmt.Println("Received:")
		fmt.Println( rmsg )
		}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
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

