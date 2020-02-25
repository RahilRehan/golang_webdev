package main

import(
	"os"
	"text/template"
)

var tpl  *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.txt"))
}

func main(){
	_ = tpl.ExecuteTemplate(os.Stdout, "file1.txt", nil)
	_ = tpl.ExecuteTemplate(os.Stdout, "file2.txt", nil)
	_ = tpl.ExecuteTemplate(os.Stdout, "file3.txt", nil)
}

