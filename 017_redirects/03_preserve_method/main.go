package main

//301 - Moves permanantly and save in cache of browser, clear cache
//303 - SeeOther, Makes a GET request to the redirected url
//307 - Preserves the Method.


import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template
var Name string

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func foo(w http.ResponseWriter, req *http.Request){
	fmt.Println("Method is: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func bar(w http.ResponseWriter, req *http.Request){
	fmt.Println("Method at bar is: ", req.Method)
	Name = req.FormValue("q")
	http.Redirect(w, req, "/barred", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request){
	fmt.Println("Method is: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", Name)
}

func main(){
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.ListenAndServe(":8080", nil)
}


