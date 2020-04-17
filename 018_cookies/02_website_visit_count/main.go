package main

import (
	"io"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("counter-cookie")
	if err == http.ErrNoCookie{
		cookie = &http.Cookie{
			Name:  "counter-cookie",
			Value: "0",
		}
	}
	count,err := strconv.Atoi(cookie.Value)
	count+=1
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)


}

func main(){
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
