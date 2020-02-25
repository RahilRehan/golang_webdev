package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func main(){
	tpl, err := template.ParseGlob("templates/*.txt")
	if err!= nil{
		log.Fatalln(err)
	}
	file, err := os.Create("templates/main.html")
	if err!= nil{
		log.Fatalln(err)
	}
	_ = tpl.ExecuteTemplate(file, "file1.txt", nil)
	_ = tpl.ExecuteTemplate(file, "file2.txt", nil)
	_ = tpl.ExecuteTemplate(file, "file3.txt", nil)

	file, _ = os.Open("templates/main.html")
	b, err := ioutil.ReadAll(file)
	fmt.Println(string(b))
}
