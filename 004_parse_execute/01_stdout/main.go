package main

import(
	"text/template"
	"log"
	"os"
)

func main(){
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err!=nil{
		log.Fatal("Error Loading File")
	}

	err = tpl.Execute(os.Stdout, nil)
	if err!=nil{
		log.Fatal("Error Loading File")
	}
}

//not optimized code, do not use