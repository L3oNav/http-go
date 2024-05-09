package main

import (
	"fmt"
	// Uncomment this block to pass the first stage
	"net"
	"os"
)
// global variable to store HTTP/1.1 200 OK\r\n\r\n
var OK_HTTP_1 = []byte("HTTP/1.1 200 OK\r\n\r\n")


func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", "0.0.0.0:4221")

  // send 200 OK
  client, err = l.Accept()
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
  client.Write(OK_HTTP_1)
  defer l.Close() 

	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}
