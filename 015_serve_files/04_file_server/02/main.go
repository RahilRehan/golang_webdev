package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,
		`<img src="/resources/toby.jpg">`,
	)
}

func main(){
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog", d)
	http.ListenAndServe(":8080", nil)
}
