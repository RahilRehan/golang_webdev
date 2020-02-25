package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main(){
	tpl, err := template.ParseFiles("file1.txt", "file2.txt")
	if err!=nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "file1.txt",nil)
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println()
	err = tpl.ExecuteTemplate(os.Stdout, "file2.txt",nil)
	if err!=nil{
		log.Fatalln(err)
	}
	fmt.Println()
	err = tpl.Execute(os.Stdout,nil)
	fmt.Println()
}
