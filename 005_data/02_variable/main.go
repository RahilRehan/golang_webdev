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
	_ = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Rahil Rehan`)
}
