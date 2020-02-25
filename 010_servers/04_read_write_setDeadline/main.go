package main

import(
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	if err!=nil{
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err!=nil{
			log.Fatal(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){

	conn.SetDeadline(time.Now().Add(10*time.Second))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I just heard you say: %s\n", ln)
	}
	defer conn.Close()
	fmt.Println("Code got here!")
}


