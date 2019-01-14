package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"homework/test_20190109_api/mongo_api_second/api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/user/all", api.GetAllUsers).Methods("GET")


	r.HandleFunc("/api/user/create", api.AddUser).Methods("GET")   //todo ID into normal represantation
	r.HandleFunc("/api/user/delete", api.DeleteUser).Methods("GET") //todo
	r.HandleFunc("/api/user/get", api.GetUser).Methods("GET")   //todo
	r.HandleFunc("/api/group/all", api.GetAllGroups).Methods("GET")  //todo
	r.HandleFunc("/api/group/get", api.GetGroup).Methods("GET")  //todo
	r.HandleFunc("/api/group/create", api.AddGroup).Methods("GET")  //todo
	r.HandleFunc("/api/group/delete", api.DeleteGroup).Methods("GET") //todo

	http.ListenAndServe(":4000", r)
}
