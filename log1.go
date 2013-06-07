package main

import (
	"os"
	"fmt"
	"log"
//	"bufio"
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

//Shaper is an interface and has a single function Area that returns an int.
type Shaper interface {
   Area() int
   Shit() int
}

type Rectangle struct {
   length, width int
}

//this function Area works on the type Rectangle and has the same function signature defined in the interface Shaper.  Therefore, Rectangle now implements the interface Shaper.
func (r Rectangle) Area() int {
   return r.length * r.width
}


func (r Rectangle) Shit() int {
	return r.length + r.width
	}

func main() {

  r := Rectangle{length:5, width:3}
   fmt.Println("Rectangle r details are: ", r)  
   fmt.Println("Rectangle r's area is: ", r.Area())  

   s := Shaper(r)
   fmt.Println("Area of the Shape r is: ", s.Area())  
   fmt.Println("Area of the Shape r is: ", s.Shit())  
   
   
   
   return;

	cfgFile := "file.txt"
	
	f, err := os.OpenFile(cfgFile, os.O_CREATE| os.O_APPEND| os.O_RDWR, 0666) // For read access.
	if err != nil {
		   checkError(err)
	}
	
  //	var l int
	 myLog := log.New(f, "PREFIXIO ", log.Lshortfile |log.Ldate|log.Ltime  );
 //	l, err = f.WriteString("HELLO THERE n");
 //	fmt.Printf("length is :%d\n", l);
 	if( err != nil ){
 		checkError2(err);
 	}
	myLog.Println("from logger");
	//myFileAppendLine(cfgFile, "HELLO THERE \n");
	
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

func myFileAppendLine ( fName string, line string )  {
	
	f, err := os.OpenFile(fName, os.O_CREATE|os.O_APPEND|os.O_WRONLY , 0644)
	checkError(err)
	/*
	if( err != nil ) {
		f, err = os.Create(fName)
		checkError(err)
		}else {
			_, err = f.Seek(0, os.SEEK_END)
		}
	*/
	//_, err = f.Seek(0, os.SEEK_END)
	_, err = f.WriteString(line)

	f.Close()
}
