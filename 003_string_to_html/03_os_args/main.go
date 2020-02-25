package main

import(
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

func main(){
	fmt.Println(os.Args[0])  //prints temp file location to run
	name := os.Args[1]
	str :=`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hello</title>
		</head>
		<body>
			<h1>Hello, World!</h1>
			<h2>` + name + `</h2>
		</body>
		</html>
	`
	file, err := os.Create("index.html")
	if err != nil{
		log.Fatal("Error creating a file!")
	}
	defer file.Close()
	fmt.Println("Writing to file")
	io.Copy(file, strings.NewReader(str))
}

// go run main.go name