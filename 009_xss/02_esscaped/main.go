package main

import (
	"log"
	"os"
	"html/template"
)

type page struct{
	Title, Heading, Input string
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main(){

	p1 := page{
		Title:   "Nothing Escaped",
		Heading: "Nothing is escaped with text/template",
		Input:   `<script>alert("Yow!");</script>`,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p1)
	if err!= nil{
		log.Fatal(err)
	}
}
