package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	if err!= nil{
		panic(err)
	}
	defer li.Close()
	for{
		conn, err := li.Accept()
		if err!= nil{
			panic(err)
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn){
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn){
	i:=0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		if i==0 && ln!=""{
			mux(conn, ln)
			i+=1
		}else{
			break
		}
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	if m=="GET" && u=="/"{
		index(conn)
	}
	if m=="GET" && u=="/about"{
		about(conn)
	}
	if m=="GET" && u=="/contact"{
		contact(conn)
	}
	if m=="GET" && u=="/apply"{
		applyProc(conn)
	}
}

func index(conn net.Conn){
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	response(conn, body)
}

func about(conn net.Conn){
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>ABOUT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	response(conn, body)
}

func contact(conn net.Conn){
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	response(conn, body)
}

func applyProc(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/about">about</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`
	response(conn, body)
}

func response(conn net.Conn, body string){
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}


