package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

type (
	Bio struct {
		Name string
		Age  uint64
	}
	User struct {
		Login    string
		Password string
		Bio
	}

	Cache map[string]User
)

var cacheByLogin Cache

func init() {
	demoUser := User{
		Login:    "userl",
		Password: "userp",
		Bio: Bio{
			Name: "usern",
			Age:  1,
		},
	}
	cacheByLogin = Cache{demoUser.Login: demoUser}
	fmt.Println(cacheByLogin)
}

func (cashMapUsers *Cache) addUser(login, pass string) {
	(*cashMapUsers)[login] = User{
		Login:    login,
		Password: pass,
	}
	//fmt.Println(*cashMapUsers, " -- into AddUser")
}

func (cashMapUsers *Cache) getUser(login string) (user User, ok bool){
	user, ok = cacheByLogin[login]
	//fmt.Println(user, " - ", ok)
	return
}

func loginCheck(login, pass string) (err error) {  // added func for adding new user to map
	if user, ok := cacheByLogin.getUser(login); ok && user.Password == pass {
		return
	}
	return errors.New("invalid pass or login")
}

func register(login, pass string) (err error) {
	fmt.Println(cacheByLogin, " - START -  at REGISTER")
	if _, ok := cacheByLogin.getUser(login); ok {
		return errors.New("user exist")
	}
	cacheByLogin.addUser(login, pass)
	fmt.Println(cacheByLogin, " - END- before return at REGISTER")
	return
}

func loginRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var login, pass string
	if _, ok := r.Form["login"]; ok {
		login = r.Form["login"][0]
	}
	if _, ok := r.Form["pass"]; ok {
		pass = r.Form["pass"][0]

	}
	err := loginCheck(login, pass)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	http.Redirect(w, r, "/?login="+login+"&pass="+pass, 303)
}

func authorizedRouterHandler(w http.ResponseWriter, r *http.Request) {
	authorized := "authorized"
	r.ParseForm()
	var login, pass string
	if _, ok := r.Form["login"]; ok {
		login = r.Form["login"][0]
	}
	if _, ok := r.Form["pass"]; ok {
		pass = r.Form["pass"][0]
	}
	err := loginCheck(login, pass)
	if err != nil {
		authorized = "unauthorized"
	}
	t, err := template.ParseFiles(
		"templates/index.html",
		"templates/"+authorized+"/header.html",
		"templates/"+authorized+"/main.html",
		"templates/footer.html",
	)

	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "index", nil)
}
func registerRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var login, pass string
	if _, ok := r.Form["login"]; ok {
		login = r.Form["login"][0]
	}
	if _, ok := r.Form["pass"]; ok {
		pass = r.Form["pass"][0]
	}
	err := register(login, pass)
	if err != nil {
		fmt.Fprint(w, "Error ")
		return
	}
	http.Redirect(w, r, "/login?login="+login+"&pass="+pass, 301)
}

func main() {

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/login", loginRouterHandler)
	http.HandleFunc("/register", registerRouterHandler)
	http.HandleFunc("/", authorizedRouterHandler)
	http.ListenAndServe(":8089", nil)
}