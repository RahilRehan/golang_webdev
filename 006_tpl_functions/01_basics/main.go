package main

import (
	"strings"
	"fmt"
	"os"
	"text/template"
)

var tpl *template.Template

//function for passing into template
func capitalize(name string) string{
	return strings.ToUpper(name)
}

func firstThree(name string) string{
	return name[:3]
}

//template funcmap to pass into template, it is map[string]interface{}
var fm = template.FuncMap{
	"uc": capitalize,
	"first3": firstThree,
}

func init(){
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.gohtml"))
}


func main(){
	p := struct{
		Name string
		Age int
	}{
		"Rahil",
		21,
	}
	fmt.Println(p)


	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p)
}
