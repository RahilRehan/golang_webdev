package main

import(
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hello, Doggy!")
}

func c(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hello, Catsy!")
}

func main(){
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", nil)
}
