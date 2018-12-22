package main

import (
	"fmt"      // пакет для форматированного ввода вывода
	"log"      // пакет для логирования
	"net/http" // пакет для поддержки HTTP протокола
	"strings"  // пакет для работы с  UTF-8 строками
)

var name, lastname, age string = "New", "User", "16"

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,

	for k, v := range r.Form {
		fmt.Println("key:", k, "val:", strings.Join(v, ""))
		if k == "name" {
			name = strings.Join(v, "")
		}
		if k == "lastname" {
			lastname = strings.Join(v, "")
		}
	}
	fmt.Fprintf(w, "Hello %s %s", name, lastname) // отправляем данные на клиентскую сторону
}

func NewRouterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //анализ аргументов,

	for k, v := range r.Form {
		fmt.Println("key:", k, "val:", strings.Join(v, ""))
		if k == "name" {
			name = strings.Join(v, "")
			fmt.Fprintf(w, "%s, сегодня тебя ждет удача! \n", name) // отправляем данные на клиентскую сторону
		}
		if k == "lastname" {
			lastname = strings.Join(v, "")
			fmt.Fprintf(w, "Твои предки по фамилии %s были поистине великими людьми\n", lastname)
		}
		if k == "age" {
			age = strings.Join(v, "")
			fmt.Fprintf(w, "О-о-о, мен, да ты так молод! Тебе только %s лет. Тебе звезды пророчат шикарное будущее \n", age)
		}
	}

}

func main() {
	http.HandleFunc("/", HomeRouterHandler) // установим роутер
	http.HandleFunc("/new", NewRouterHandler)
	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
