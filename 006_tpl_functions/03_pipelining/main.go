package main

import(
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func double(num int) int {
	return num*2
}

func square(num int) float64{
	return math.Pow(float64(num), 2)
}

func sqRoot(num float64) float64{
	return math.Sqrt(num)
}

var fm = template.FuncMap{
	"d" : double,
	"sq" : square,
	"sqr" : sqRoot,
}

func init(){
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main(){
	_ = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 4)
}