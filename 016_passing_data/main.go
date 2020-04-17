package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"net/http"
)

var tpl *template.Template

type person struct {
	Name, Age, FileContent string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func store(w http.ResponseWriter, req *http.Request) {

	var p person
	if req.Method == http.MethodPost {
		//read file
		f, h, err := req.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer f.Close()

		fmt.Println("\nFile", f, "\nHeader", h, "\nError", err)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s := string(bs)

		//store file
		dst, err := os.Create(filepath.Join("./user/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		p = person{
			req.FormValue("name"),
			req.FormValue("age"),
			s,
		}

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.gohtml", p)
}

func main() {
	http.HandleFunc("/", store)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
