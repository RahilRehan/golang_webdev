package main

import (
	"bufio"
	"fmt"
	"net"
)

func main(){
	li, err := net.Listen("tcp", "localhost:8080")
	if err!=nil{
		panic(err)
	}

	defer li.Close()

	for{
		conn, err := li.Accept()
		if err!=nil{
			panic(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == ""{
			break
		}
	}
}

func response(conn net.Conn){
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}