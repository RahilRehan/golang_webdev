package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
	ages := map[string]int{
		"Rahil" : 21,
		"Bhanu": 20,
	}

	_ = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", ages)
}
