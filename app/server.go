package main

import (
  "bufio"
	"fmt"
  "net"
	"os"
  "strings"
  "bytes"
)

type HTTPRequest struct {
	Method    string
	Path      string
	Headers   map[string]string
	Body      string
	UserAgent string
}

// global variable to store HTTP/1.1 200 OK\r\n\r\n
var OK_HTTP_1 = []byte("HTTP/1.1 200 OK\r\n\r\n")

func getStatus(statusCode int, statusText string) string {
	return fmt.Sprintf("HTTP/1.1 %d %s", statusCode, statusText)
}

func parseStatus(scanner *bufio.Scanner) (*HTTPRequest, error) {
	var req HTTPRequest = HTTPRequest{}
	req.Headers = make(map[string]string)
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			parts := strings.Split(scanner.Text(), " ")
			req.Method = parts[0]
			req.Path = parts[1]
			continue
		}
		headers := strings.Split(scanner.Text(), ": ")
		if len(headers) < 2 {
			req.Body = headers[0]
			break
		}
		if headers[0] == "User-Agent" {
			req.UserAgent = headers[1]
		}
		req.Headers[headers[0]] = headers[1]
	}
	return &req, nil
}

func Handler(conn net.Conn) {
  defer conn.Close()
  // read the request

  scanner := bufio.NewScanner(conn)

  req, err := parseStatus(scanner)

  if err != nil {
    fmt.Println("Error reading request: ", err.Error())
    return
  }
 
  var response string

  switch path := req.Path; {
	  case strings.HasPrefix(path, "/echo/"):
		  content := strings.TrimLeft(path, "/echo/")
		  response = fmt.Sprintf("%s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", getStatus(200, "OK"), len(content), content)
	  case path == "/user-agent":
		  response = fmt.Sprintf("%s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", getStatus(200, "OK"), len(req.UserAgent), req.UserAgent)
	  case path == "/":
		  response = getStatus(200, "OK") + "\r\n\r\n"
    case strings.HasPrefix(path, "/files/") && req.Method == "GET":
      dir := os.Args[2]
      fileName := strings.TrimPrefix(path, "/files/")
      fmt.Println("dir: ", dir, " fileName: ", fileName)
      data, error := os.ReadFile(dir + "/" + fileName)

      if error != nil {
        fmt.Println("Error reading file: ", error.Error(), " for file: ", fileName, " in directory: ", dir)
        response = getStatus(404, "Not Found") + "\r\n\r\n"
      } else {
        response = fmt.Sprintf("%s\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", getStatus(200, "OK"), len(data), data)
      }
    case strings.HasPrefix(path, "/files/") && req.Method == "POST":
      dir := os.Args[2]
      fileName := strings.TrimPrefix(path, "/files/")
      fmt.Println("body: ", req.Body)
      contents := bytes.Trim([]byte(req.Body), "\x00")
      err := os.WriteFile(dir + fileName, contents, 0644)
      if err != nil {
        fmt.Println("Error writing file: ", err.Error())
        response = getStatus(500, "Internal Server Error") + "\r\n\r\n"
      } else {
        response = getStatus(200, "OK") + "\r\n\r\n"
      }
	  default:
		  response = getStatus(404, "Not Found") + "\r\n\r\n"
	}

  conn.Write([]byte(response))
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

  for {
    conn, err := l.Accept()
    if err != nil {
      fmt.Println("Error accepting: ", err.Error())
      os.Exit(1)
    }

    go Handler(conn)
  }
}



