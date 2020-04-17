package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/set", setRoute)
	http.HandleFunc("/read", readRoute)
	http.HandleFunc("/expire", expireRoute)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, `<h1><a href="/set">SET COOKIE</a></h1>`)
}

func setRoute(res http.ResponseWriter, req *http.Request){
	http.SetCookie(res, &http.Cookie{
		Name:  "session",
		Value: "Some value",
	})
	fmt.Fprintf(res, `<h1><a href="/read">READ COOKIE</a></h1>`)
}

func readRoute(res http.ResponseWriter, req *http.Request){
	c, err := req.Cookie("session")
	if err != nil{
		http.Redirect(res, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(res, `<h1>%v</h1><br><h1><a href="/expire">EXPIRE</a></h1>`, c)
}

func expireRoute(res http.ResponseWriter, req *http.Request){
	c, err := req.Cookie("session")
	if err != nil{
		http.Redirect(res, req, "/set", http.StatusSeeOther)
	}
	c.MaxAge = -1
	http.SetCookie(res, c)
	http.Redirect(res, req, "/", http.StatusSeeOther)
}