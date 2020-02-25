package main

import(
	"bufio"
	"fmt"
	"net"
)

func main(){
	li, err := net.Listen("tcp", "localhost:8080")
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
		go handle(conn)
	}
}

func handle(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		bs := []byte(ln)
		out := rot13(bs)
		fmt.Fprintf(conn, "Rot13: %s\n", string(out))
	}
	fmt.Println("Code Reached Here!")
	defer conn.Close()
}

func rot13(bs []byte) []byte{
	out := make([]byte, len(bs))
	for i, v := range bs{
		if v<=109 {
			out[i] = v+13
		}else{
			out[i] = v-13
		}
	}
	return out
}