package main

import (
	"fmt"
  "net"
  "bufio"
	"net/http"
	"os"
)
// global variable to store HTTP/1.1 200 OK\r\n\r\n
var OK_HTTP_1 = []byte("HTTP/1.1 200 OK\r\n\r\n")

func Handler(conn net.Conn) {
  defer conn.Close()
  // read the request
  request, err := http.ReadRequest(bufio.NewReader(conn))
  if err != nil {
    fmt.Println("Error reading request: ", err.Error())
    return
  }
  if request.URL.Path == "/" {
    conn.Write(OK_HTTP_1)
    return
  }
  conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
}


func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
  l, err := net.Listen("tcp", ":4221")
  if err != nil {
    fmt.Println("Error listening: ", err.Error())
    os.Exit(1)
  }

  defer l.Close()

  conn, err := l.Accept()
  if err != nil {
    fmt.Println("Error accepting: ", err.Error())
    os.Exit(1)
  }

  Handler(conn)

}



