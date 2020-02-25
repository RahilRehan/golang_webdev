package main

import(
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil{
		log.Fatalln(err)
	}
}