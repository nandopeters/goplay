package main

import (
	"fmt"
//	"io"
	"os"
//	"bytes"
	)
	
		


func main() {
	input_len := 100
	var input []byte = make( []byte, input_len)


	for n := 0 ; n < 2; n++ {
	// read from stdin
	cnt, err:= os.Stdin.Read( input)

	if err != nil {
			fmt.Printf("error:%d\n", err)
	} 

	myStr := string( rTrim(input, cnt-1) )
	fmt.Println( myStr)	
	}
}

func rTrim( input []byte, cnt int) ([]byte){
	for i:= cnt; i < len(input) ; i++ {
		input[i] = 0
	} 
	return input
}


