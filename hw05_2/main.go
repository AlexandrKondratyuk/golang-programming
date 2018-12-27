package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/logon", LogOnHandler)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	// fmt.Println(userMain)
}

//HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, "Hello, User!")
}

// LogOnHandler ...
func LogOnHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	for k, v := range r.Form {
		var vNew string = strings.Join(v, "")
		switch k {
		case "login":
			fmt.Fprintf(w, "%T = %v", vNew, vNew)
		case "name":
			fmt.Fprintf(w, "%T = %v", vNew, vNew)
		case "lastname":
			fmt.Fprintf(w, "%T = %v", vNew, vNew)
		case "age":
			fmt.Fprintf(w, "%T = %v", vNew, vNew)
		}
	}

	fmt.Fprintln(w, "LogOn!")
}
