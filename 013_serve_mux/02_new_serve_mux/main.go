package main

import(
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hotdog")
}

func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hotcat")
}

func main(){
	var d hotdog
	var c hotcat
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)
	http.ListenAndServe(":8080", mux)
}

/*
/item/ works as /item as well as /item/something in NewServerMux() type
404 not found is implicitly defined and handled by  NewServerMux() type
 */
