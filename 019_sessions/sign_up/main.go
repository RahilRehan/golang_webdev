package  main

import(
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct{
	Username string
	Password string
}

var dbUsers = map[string]user{}   //username, user struct
var dbSessions = map[string]string{} //session id(Sid), username

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/signUp", signUp)
	http.HandleFunc("/logOut", logOut)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request){
	if !LoggedIn(req){
		tpl.ExecuteTemplate(res, "signup.gohtml", nil)
		return
	}
	c, err := req.Cookie("session")
	if err!=nil{
		http.Redirect(res, req, "/", http.StatusForbidden)
	}
	un := dbSessions[c.Value]
	u := dbUsers[un]
	tpl.ExecuteTemplate(res, "show.gohtml", u)
}

func signUp(res http.ResponseWriter, req *http.Request){

	if !LoggedIn(req){
		un := req.FormValue("username")
		p := req.FormValue("password")
		fmt.Println(un, p)

		if _, ok := dbUsers[un]; ok{
			fmt.Fprintf(res, "User with username already exists")
			return
		}
		sId, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(res, c)
		dbSessions[sId.String()] = un
		u := user{un, p}
		dbUsers[un] = u

		tpl.ExecuteTemplate(res, "show.gohtml", u)
	}

	http.Redirect(res, req, "/", http.StatusSeeOther)

}

func logOut(res http.ResponseWriter, req *http.Request){
	if LoggedIn(req){
		c, _ := req.Cookie("session")
		c.MaxAge = -1
		http.SetCookie(res, c)
		tpl.ExecuteTemplate(res, "show.gohtml", nil)
	}
}

func LoggedIn(req *http.Request) bool {
	_,err := req.Cookie("session")
	if err!=nil{
		return false
	}
	return true
}