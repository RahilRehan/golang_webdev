package main

import (
	"html/template"
	"log"
	"os"
)

type person struct{
	Name string
	Age int
}

func (p person) DblAge() int{
	return 2*p.Age
}


var tpl *template.Template

func init(){
	tpl = template.Must(template.New("index.gohtml").Funcs(template.FuncMap{
		"realAge" : func(age int) int{
		return age/2
	},
	}).ParseFiles("index.gohtml"))
}

func main(){
	p1 := person{
		"Rahil",
		19,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p1)
	if err!=nil{
		log.Fatal(err)
	}
}
