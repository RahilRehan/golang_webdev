package main

import(
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Hello, People!</h1>`)
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080", d)
}
