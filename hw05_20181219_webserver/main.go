package main

import (
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	"strings"  // пакет для работы с  UTF-8 строками
)

func main() {
	http.HandleFunc("/", HomeRouterHandler) // установим роутер
	http.HandleFunc("/new", NewRouterHandler)
	err := http.ListenAndServe(":3000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

var name, lastname, age string = "New", "User", "16"

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) { // дефолтное view
	r.ParseForm() //анализ аргументов,

	for k, v := range r.Form {

		switch k {
		case "name":
			name = strings.Join(v, "")
		case "lastname":
			lastname = strings.Join(v, "")
		}
	}
	fmt.Fprintf(w, "Hello %s %s", name, lastname) // отправляем данные на клиентскую сторону
}

func NewRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,

	for k, v := range r.Form {
		fmt.Println("key:", k, "val:", strings.Join(v, ""))

		switch k {
		case "name":
			name = strings.Join(v, "")
			fmt.Fprintf(w, "Your name is %s! \n", name)
		case "lastname":
			lastname = strings.Join(v, "")
			fmt.Fprintf(w, "Your surname is  %s \n", lastname)
		case "age":
			age = strings.Join(v, "")
			fmt.Fprintf(w, "You are %s years old \n", age)
		}
	}
}
