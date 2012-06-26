package main

import "fmt"
import "os"

func main() {
  var input []byte = make( []byte, 100 );
  os.Stdin.Read( input );
  fmt.Printf( "%s", input );
}

