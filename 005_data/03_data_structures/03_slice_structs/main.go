package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

type person struct{
	Name string
	Age int
}

func main(){

	bhanu := person{"Govardhan", 19}
	rahil := person{"Rahil", 21}

	users := []person{bhanu, rahil}

	_ = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
}
