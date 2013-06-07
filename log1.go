package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
)



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










func main() {

	w:= bufio.NewWriter(os.Stdout)
	w.WriteString("HELLO THERE \nOK\n\n");

	fmt.Printf("Flags:%q   Prefix:%q\n", log.Flags(), log.Prefix() );
	os.Exit(0);
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

