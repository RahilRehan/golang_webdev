package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main(){
	_ = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 54)
}