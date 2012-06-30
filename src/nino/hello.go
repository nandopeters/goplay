package main

import (
	"fmt"
	"encoding/json"
	"os"
	"time"
	)
	
type Location struct {
	Ag		string
	Sb		string	
	DT		string

}

type BCHAN struct {
	q map[string]chan string
	}
func  (cm *BCHAN) initQ ( ) {
	cm.q = make(map[string]chan string)
	}

func  (cm BCHAN) addQ ( who string) {
	cm.q[who] =  make(chan string) 
	}

func  (cm *BCHAN) popQ ( who string ) string {
	return <- cm.q[who]
}

func  (cm *BCHAN) popAllQ ( who string )  {
		for s := range cm.q[who] {
			fmt.Println("pop:",s)
		}
	}

type ACHAN struct {
	c		chan string
	}
type CHAN2 struct {
	q map[string]ACHAN
	}
	
var	ac ACHAN


func doit (c chan string )  {
	c <- "msg 1"
	c <- "msg 2"
	}
func  (cm *CHAN2) initQ2 ( ) {
	cm.q = make(map[string]ACHAN)
	}
func  (cm *CHAN2) popQ2 ( who string ) string {
	return <- cm.q[who].c
	}
func  (cm *CHAN2) popAllQ2 ( who string )  {
		for s := range cm.q[who].c {
			fmt.Println("pop:",s)
		}
	}
func  (cm CHAN2) addQ2 ( who string) {
	cm.q[who] = ACHAN{  make(chan string) }
	}
	
func  addQ (  q map[string]ACHAN, who string) {
	q[who] = ACHAN{  make(chan string) }
	}
func setQ ( q map[string]ACHAN, who string) chan string {
	return q[who].c
}
func main() {

var b BCHAN

b.initQ()
b.addQ("AL")
b.addQ("BOB")
go doit( b.q["AL"])
go doit( b.q["BOB"])


fmt.Println(b.popQ("BOB"))
fmt.Println(b.popQ("BOB"))
fmt.Println(b)

os.Exit(0)


var c CHAN2
c.initQ2()
c.addQ2("AL")
c.addQ2("BOB")
go doit( c.q["AL"].c)
go doit( c.q["BOB"].c)

//fmt.Println(c.popQ2("AL"));
//fmt.Println(c.popQ2("AL"))
//c.addQ2("AL")

fmt.Println(c.popQ2("BOB"))
fmt.Println(c.popQ2("BOB"))
fmt.Println(c)

 os.Exit(0)

// Make a map of type []ACHAN
/*
achan := make ( map[string] []ACHAN)
achan["AL"] = append(achan["AL"], ac )
fmt.Println(achan)
*/

// Make a map of type []string
/*
mm := make( map[string][]string ) 
mm["AL"] = append(mm["AL"],"")
pop(mm["AL"])
mm["AL"] = append(mm["AL"], "message 1")
mm["BOB"] = push(mm["BOB"], "Bob Message 111")


var ss string
ss, mm["AL"] = pop(mm["AL"])
fmt.Println(mm["AL"])
fmt.Println(ss)
*/




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
	
func rTrim( input []byte, cnt int) ([]byte){
	for i:= cnt; i < len(input) ; i++ {
		input[i] = 0
	} 
	return input
}


