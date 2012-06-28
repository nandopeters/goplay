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
	DT		string

}

func str2js(c Location) string {
	cjs, err := json.Marshal(c)	
	if err != nil {
        return "Json failed"
    }
    return string(cjs)
}

func locationUpdate() string {
	
	loc := Location{}
	loc.Ag = "locmgr"
	loc.Sb = "locupdt"
	loc.DT = time.Now().String()
	jstr := str2js(loc)   
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


	
func pop ( s [] string ) (string, []string ) {
	l:= len(s)
	if ( l == 0 ){
		return "", s[0:0]
		}
	return s[0], s[1: l]
}

func push ( s [] string, item string) []string {
	return append (s, item)
	}
	
type CS struct {
	name	string
	msg 	[] string

}

type CM struct {
	 cs 	*CS
	}

var m map[string]CM

func main() {
cs := new (CS)


cs.name = "AL"
cs.msg = push (cs.msg, "'message1 from AL'" )
cs.msg = push (cs.msg, "'message2 from AL'" )


m = make ( map[string]CM)
m["AL"] = CM { cs }
//fmt.Println( m["AL"].cs )


mm := make( map[string][]string ) 
fmt.Println("length of map", len(mm))
mm["AL"] = append(mm["AL"],"")
pop(mm["AL"])
fmt.Println( mm )

mm["AL"] = append(mm["AL"], "message 1")
mm["AL"] = append(mm["AL"], "message 2")
mm["AL"] = append(mm["AL"], "message 3")
mm["BOB"] = push(mm["BOB"], "Bob Message 111")
mm["BOB"] = push(mm["BOB"], "Bob Message 222")
fmt.Println("length of map", len(mm))
for k, _ := range mm {
	fmt.Println(k)
}
fmt.Println( len(mm["AL"]) )
os.Exit(0)
var ss string
ss, mm["AL"] = pop(mm["AL"])
fmt.Println(mm["AL"])
fmt.Println(ss)
//var  cm   CM
//cm.cs = cs;
//fmt.Println(cm)



}

func rTrim( input []byte, cnt int) ([]byte){
	for i:= cnt; i < len(input) ; i++ {
		input[i] = 0
	} 
	return input
}


