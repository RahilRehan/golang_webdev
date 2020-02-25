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

type car struct{
	RaceCar bool
	Milage int
}


func main(){

	bhanu := person{"Govardhan", 19}
	rahil := person{"Rahil", 21}

	ferrari := car{true, 7}
	wagnor := car{false, 15}

	users := []person{bhanu, rahil}
	cars := []car{ferrari, wagnor}

	data := struct {
		People []person
		Vehicles []car
	}{
		users,
		cars,
	}
	_ = tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
}
