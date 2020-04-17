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
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))
	http.ListenAndServe(":8080", nil)
}
