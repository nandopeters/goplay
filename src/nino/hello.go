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

func main() {

service := "part1 : "

service += "part2"

fmt.Println( service )
os.Exit(0)

fmt.Println( time.Now().String() )

readALine();


smsg := locationUpdate()
fmt.Println( smsg)	

}

func rTrim( input []byte, cnt int) ([]byte){
	for i:= cnt; i < len(input) ; i++ {
		input[i] = 0
	} 
	return input
}


