package main

import(
	"fmt"
)

func main(){
	name := "Rahil Rehan"
	text := `
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
	fmt.Println(text)
}