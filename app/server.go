package main

import (
	"fmt"
  "net"
	"os"
  "strings"
)
// global variable to store HTTP/1.1 200 OK\r\n\r\n
var OK_HTTP_1 = []byte("HTTP/1.1 200 OK\r\n\r\n")

func Handler(conn net.Conn) {
  defer conn.Close()
  // read the request

  buff := make([]byte, 1024)

  request, err := conn.Read(buff)
  path := strings.Split(string(buff[:request]), " ")[1]

  if err != nil {
    fmt.Println("Error reading request: ", err.Error())
    return
  }

  if path == "/" {
    conn.Write(OK_HTTP_1)
    return
  } else if strings.Split(path, "/")[1] == "echo" {
    message := strings.Split(path, "/")[2]
    conn.Write([]byte(fmt.Sprintf(
      "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(message), message)     ))
  } else if path == "/user" {
    conn.Write([]byte(fmt.Sprintf("%s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", getStatus(200, "OK"), len(req.UserAgent), req.UserAgent))
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



