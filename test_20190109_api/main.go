package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type User struct {
	ID    string `json:"id,omitempty"`
	Login string `json:"login,omitempty"`
	Pass  string `json:"pass,omitempty"`
	Age   int    `json:"age,omitempty"`
}

//type Address struct {
//	City  string `json:"city,omitempty"`
//	State string `json:"state,omitempty"`
//}

var Users []User

func get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range Users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}
func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}
func create(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)  //// Vars returns the route variables for the current request, if any. return rv.(map[string]string)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	Users = append(Users, user)
	json.NewEncoder(w).Encode(Users)
}
func delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range Users {
		if item.ID == params["id"] {
			Users = append(Users[:index], Users[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(Users)
	}
}

func main() {
	router := mux.NewRouter() // определяем маршруты с помощью gorilla/mux
	Users = append(Users, User{ID: "1", Login: "Alex", Pass: "Kon", Age: 32})
	Users = append(Users, User{ID: "2", Login: "Peter", Pass: "Block", Age: 27})
	router.HandleFunc("/api/user/getall", getAll).Methods("GET")
	router.HandleFunc("/api/user/get/{id}", get).Methods("GET")
	router.HandleFunc("/api/user/create/{login, pass, age}", create).Methods("GET")
	router.HandleFunc("/api/user/delete/{id}", delete).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", router))
}
