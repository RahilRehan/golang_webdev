package main

import(
	"fmt"
	"net/http"
)

type hotdog int

func (f hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w, "Hey, Maamulgundadu!")
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080", d)
}
