package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(bs))
}

/*
run 01_write/main.go and also 05_dial_read.go -> reads from the connection by dialing into the connection
*/