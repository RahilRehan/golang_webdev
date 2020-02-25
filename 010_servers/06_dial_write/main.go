package main

import (
"fmt"
"log"
"net"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "Hey, I dialed you!")
}

/*
run 02_read/main.go and also 06_dial_write.go
*/



