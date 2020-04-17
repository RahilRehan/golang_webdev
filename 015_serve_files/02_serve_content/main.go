package main

import(
	"io"
	"log"
	"net/http"
	"os"
)

func d(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,
		`First Image:
			<img src="https://cdn.pixabay.com/photo/2016/12/13/05/15/puppy-1903313__340.jpg"><br>
			Second Image:
			<img src="/toby.jpg">`,
	)
}

func dogPic(w http.ResponseWriter, req *http.Request){
	file, err := os.Open("toby.jpg")
	if err!=nil{
		http.Error(w, "File not found", 404)
		log.Fatal(err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err!=nil{
		log.Fatal(err)
	}
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}

func main(){
	http.HandleFunc("/dog", d)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}
