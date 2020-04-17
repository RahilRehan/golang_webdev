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

func init(){
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func foo(w http.ResponseWriter, req *http.Request){
	fmt.Println("Method is: ", req.Method)
	w.Header().Set("Location", "/bar")
	w.WriteHeader(http.StatusMovedPermanently)
}

func bar(w http.ResponseWriter, req *http.Request){
	fmt.Println("Method at bar is: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func main(){
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}


