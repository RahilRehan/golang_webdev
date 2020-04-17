package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,
		`<img src="toby.jpg">`,
	)
}

func main(){
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", d)
	http.ListenAndServe(":8080", nil)
}

//problem , serves all the files when navigated to index
