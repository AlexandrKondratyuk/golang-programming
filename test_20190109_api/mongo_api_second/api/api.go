package api

import (
	"net/http"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
	"homework/test_20190109_api/mongo_api_second/db"
)

func handleError(err error, message string, w http.ResponseWriter) { //handle errors for all methods
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

func GetAllUsers(w http.ResponseWriter, req *http.Request) { // returns a list of all users to response
	rs, err := db.GetAllUsers()
	if err != nil {
		handleError(err, "Failed to load database users: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}
	w.Write(bs)
}

func AddUser(w http.ResponseWriter, req *http.Request) { // Add s user into the database users
	login := req.FormValue("login")
	pass := req.FormValue("pass")
	ageString := req.FormValue("age")
	age, err := strconv.Atoi(ageString)

	if err != nil {
		handleError(err, "Failed to parse input data: %v", w)
		return
	}

	user := db.User{Login: login, Pass: pass, Age: age}

	if err = db.SaveUser(user); err != nil {
		handleError(err, "Failed to save user: %v into db users", w)
		return
	}

	res, _ := json.Marshal(user.Login)
	//res := user.Login

	w.Write([]byte(res))
}

func DeleteUser(w http.ResponseWriter, req *http.Request) { // Removes user by ID from the db users
	req.ParseForm()
	var login string

	if _, ok := req.Form["login"]; ok {
		login = req.Form["login"][0]
	}

	fmt.Println("id for delete = ", login)

	if err := db.RemoveUser(login); err != nil {
		handleError(err, "Failed to remove user: %v\n", w)
		w.Write([]byte("OK: false"))
		return
	}

	w.Write([]byte("OK: true"))
}

func GetUser(w http.ResponseWriter, req *http.Request) { // get user by idfrom DB users
	req.ParseForm()
	var login string

	if _, ok := req.Form["login"]; ok {
		login = req.Form["login"][0]
	}

	rs, err := db.GetOneUser(login)

	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}
	w.Write(bs)
}

func FindLogin(w http.ResponseWriter, req *http.Request) { // get user by id from DB users
	req.ParseForm()
	var login string
	if _, ok := req.Form["login"]; ok {
		login = req.Form["login"][0]
	}

	rs, err := db.FindIdByLogin(login)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs.Login)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}
	w.Write(bs)
}





////////////////////
// TODO done to here

func GetAllGroups(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAllGroups()
	if err != nil {
		handleError(err, "Failed to load database groups: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}

func GetGroup(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOneGroup(id)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}

func AddGroup(w http.ResponseWriter, req *http.Request) {
	ID := req.FormValue("id")
	valueStr := req.FormValue("value")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		handleError(err, "Failed to parse input data: %v", w)
		return
	}

	group := db.Group{ID: ID, Value: value}

	if err = db.SaveGroup(group); err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}

	w.Write([]byte("OK"))
}

func DeleteGroup(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := db.RemoveGroup(id); err != nil {
		handleError(err, "Failed to remove item: %v", w)
		return
	}

	w.Write([]byte("OK"))
}
