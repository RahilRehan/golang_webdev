package main

import(
	"os"
	"log"
	"io"
	"strings"
	"fmt"
)

func main(){
	str :=`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hello</title>
		</head>
		<body>
			<p>Hello, World!</p>
		</body>
		</html>
	`
	file, err := os.Create("index.html")
	if err!= nil{
		log.Fatal("Error creating a File!")
	}

	defer file.Close()
	fmt.Println("Writing to File!")
	io.Copy(file, strings.NewReader(str))

}