package api

import (
	"net/http"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"strconv"
	"homework/test_20190109_api/mongo_api_second/db"
	"gopkg.in/mgo.v2/bson"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// GetAllItems returns a list of all database items to the response.
func GetAllUsers(w http.ResponseWriter, req *http.Request){
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

// Add one user (form data) into the database
func AddUser(w http.ResponseWriter, req *http.Request) {
	ID := req.FormValue("id")
	//ID := bson.NewObjectId()
	//ID := string(bson.NewObjectId())
	login := req.FormValue("login")
	pass := req.FormValue("pass")
	ageString := req.FormValue("age")
	age, err := strconv.Atoi(ageString)
	if err != nil {
		handleError(err, "Failed to parse input data: %v", w)
		return
	}

	user := db.User{ID: ID, Login: login, Pass: pass, Age: age}

	if err = db.SaveUser(user); err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}

	bs, _ := json.Marshal(user.ID)
	w.Write([]byte(bs)) //TODO return json formatted ID
}

// Removes a single item (identified by id) from the database.
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := bson.ObjectId(vars["id"])

	if err := db.RemoveUser(id); err != nil {
		handleError(err, "Failed to remove user: %v", w)
		w.Write([]byte("OK: false"))
		return
	}

	w.Write([]byte("OK: true"))
}

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

// GetItem returns a single database item matching given ID parameter.
func GetUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOneUser(id)
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