package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	username string
	password string
}

var Users = make(map[string]User)

func main() {
	Users["admin"] = User{"admin", "pass"} // add default admin's data into map

	http.HandleFunc("/welcome", WelcomePage) // setting router rules
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logon", Logon)
	http.HandleFunc("/", Home)

	err := http.ListenAndServe(":9090", nil) // setting listening port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("End MAIN")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start HOME")
	fmt.Println("==>", Users)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("hello.html") // return (*Template, error)
		t.Execute(w, nil)

		r.ParseForm()

		if len(r.Form) > 0 {
			_, ok1 := Users[r.Form["username"][0]]
			_, ok2 := Users[r.Form["password"][0]]
			if ok1 && ok2 {
				http.Redirect(w, r, "/welcome", http.StatusFound)
			}
		}
	}

	fmt.Println("path", r.URL.Path)
	fmt.Println("******************************End HOME")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start logIN")
	fmt.Println("path", r.URL.Path)
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		passFromMap, ok := Users[r.Form["username"][0]]
		passFromForm := Users[r.Form["password"][0]]
		if ok && passFromForm == passFromMap {
			http.Redirect(w, r, "/welcome", http.StatusFound)
		} else {
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
	fmt.Println("******************************End logIN")
}

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("welcome.html") // return (*Template, error)
	t.Execute(w, nil)
}

func Logon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Start logON")
	fmt.Println("path", r.URL.Path)
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("logon.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		_, ok := Users[r.Form["username"][0]]
		//passFromMap, ok := Users[r.Form["username"][0]]
		//passFromForm := Users[r.Form["password"][0]]

		if ok {
			fmt.Fprintf(w, `<html>
            <head>
            </head>
            <body>
			<script>
				alert('This login is already in use. Try another!')
			</script>            
            </script>
            </body>
            </html>`)

			t, _ := template.ParseFiles("logon.html")
			t.Execute(w, nil)
		} else {
			r.ParseForm()
			login, pass := r.Form["username"][0],r.Form["password"][0]
			Users[login] = User{login, pass}
			fmt.Println("New Users map ==>", Users)

			fmt.Fprintf(w, `<html>
            <head>
            </head>
            <body>
			<script>
				alert('Your data saved seccessfully!')
			</script>            
            </script>
            </body>
            </html>`)

			t, _ := template.ParseFiles("login.html")
			t.Execute(w, nil)
		}
	}
	fmt.Println("******************************End LOGIN")
}
