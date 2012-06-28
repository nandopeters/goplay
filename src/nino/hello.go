package main

import (
	"fmt"
	"encoding/json"
//	"io"
	"os"
//	"bytes"
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
	fmt.Println( loc )
	jstr := str2js(loc)   
	//jstr += "}"
	
	return jstr
}

func readALine()  {
	input_len := 100
	var input []byte = make( []byte, input_len)
	for n := 0 ; n < 4; n++ {
	// read from stdin
	cnt, err:= os.Stdin.Read( input)

	if err != nil {
			fmt.Printf("error:%d\n", err)
	} 

	myStr := string( rTrim(input, cnt-1) )
	fmt.Println( "line : ", n, myStr)	
	}
}

type CS struct {
	name	string
	msg 	[] string

}
	
func pop ( s [] string ) (string, []string ) {
	
	l:= len(s)
	return s[0], s[1: l]

}

func push ( s [] string, item string) []string {
	return append (s, item)
	}
	


func main() {


cs := new (CS)

var s [] string
var r string
s = append ( s, "Msg 1" )
s = push ( s, "Msg 2")
s = push (s, "Msg 3")
s = push (s, "Msg 4")

fmt.Println ( len(s) )

r , s = pop ( s)
fmt.Println ( r );
fmt.Println ( s );




}

func rTrim( input []byte, cnt int) ([]byte){
	for i:= cnt; i < len(input) ; i++ {
		input[i] = 0
	} 
	return input
}


