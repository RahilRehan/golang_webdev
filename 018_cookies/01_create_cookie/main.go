package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func addcookies(w http.ResponseWriter, req *http.Request){
	http.SetCookie(w, &http.Cookie{
		Name:       "Loner",
		Value:      "I am loner cookie",
	})
	//http.SetCookie(w, &http.Cookie{
	//	Name:       "Counter!",
	//	Value:      "I am loner cookie",
	//})
}

func readCookie(w http.ResponseWriter, req *http.Request){
	c, err := req.Cookie("Loner")
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(c)
	s := fmt.Sprintf("Cookie's name is %s and value is %s: ", c.Name, c.Value)
	io.WriteString(w, s)
}

func main(){
	http.HandleFunc("/", addcookies)
	http.HandleFunc("/read", readCookie)
	http.ListenAndServe(":8080", nil)
}
