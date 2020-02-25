package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
}
