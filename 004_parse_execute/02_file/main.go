package main

import(
	"log"
	"os"
	"text/template"
)

func main(){
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err!=nil{
		log.Fatalln(err)
	}

	file, err := os.Create("index.html")
	if err!=nil{
		log.Fatalln(err)
	}
	defer file.Close()

	err = tpl.Execute(file, nil)
	if err!=nil{
		log.Fatalln(err)
	}	
}