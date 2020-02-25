package main

import(
	"os"
	"text/template"
)

var tmp *template.Template

func init(){
	tmp = template.Must(template.ParseFiles("tpl.gohtml"))
}

type user struct{
	Name string
	Age int
	Check bool
	Ar []int
}

func main(){
	u1 := user{
		"Rahil",
		21,
		false,
		[]int{},
	}
	u2 := user{
		"",
		20,
		true,
		[]int{},
	}
	u3 := user{"Rehan",
		17,
		true,
		[]int{3,2,1},
	}
	users := []user{u1, u2, u3}

	_ = tmp.ExecuteTemplate(os.Stdout, "tpl.gohtml", users)
}
