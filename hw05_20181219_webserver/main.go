package main

import (
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	"strings"  // пакет для работы с  UTF-8 строками
)

var name, lastname, age string = "New", "User", "16"

type User struct {
	Name     string
	Lastname string
	Age      string
}

func main() {
	userMain := &User{"undefined Name", "undefined LastNane", "undefined Age"}
	// myCash := new(Cash)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintf(w, "Hello %s %s", name, lastname)
	})

	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm() //анализ аргументов,

		for k, v := range r.Form {
			switch k {
			case "name":
				userMain.setName(strings.Join(v, ""))
			case "lastname":
				userMain.setLastName(strings.Join(v, ""))
			case "age":
				userMain.setAge(strings.Join(v, ""))
			}
		}
	})
	http.HandleFunc("/logon", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm() //анализ аргументов,

		for k, v := range r.Form {
			switch k {
			case "name":
				userMain.setName(strings.Join(v, ""))
			case "lastname":
				userMain.setLastName(strings.Join(v, ""))
			case "age":
				userMain.setAge(strings.Join(v, ""))
			}
		}

		fmt.Fprintf(w, "You registered name => %s! \n", userMain.Name)
		fmt.Fprintf(w, "You registered lastname => %s \n", userMain.Lastname)
		fmt.Fprintf(w, "You registered age => %s \n", userMain.Age)
	})

	err := http.ListenAndServe(":3000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println(userMain)
}

func (user *User) setName(name string) {
	user.Name = name
}

func (user *User) setLastName(lastname string) {
	user.Lastname = lastname
}

func (user *User) setAge(age string) {
	user.Age = age
}

type Cash struct {
	Users map[string]User
}

var mapUsers map[string]User = make(map[string]User)

func (cash *Cash) SetCash(login string, user *User) {
	(*cash).Users[login] = *user
}
