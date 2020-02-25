package main

import(
	"html/template"
	"log"
	"os"
)

type person struct{
	Name string
	Age int
}

type hero struct{
	person
	Ltokill bool
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
	p1 := hero{
		person{
			"Rahil",
			19,
		},
		false,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err!= nil{
		log.Fatalln(err)
	}
}
