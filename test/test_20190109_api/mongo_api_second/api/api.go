package api

import (
	"net/http"

	"encoding/json"
	"fmt"
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

// GROUPS

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

func AddGroup(w http.ResponseWriter, req *http.Request) {
	groupName := req.FormValue("groupName")
	fmt.Println("groupName => ", groupName)

	users := []db.User{}
	fmt.Println("users => ", users)

	group := db.Group{GroupName: groupName, Users: users}
	fmt.Println("group => ", group)

	if err := db.SaveGroup(group); err != nil {
		handleError(err, "Failed to save group: %v into db groups", w)
		return
	}

	res, _ := json.Marshal(group.GroupName)
	fmt.Println("res", res)
	w.Write([]byte(res))
}

func DeleteGroup(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var groupName string

	if _, ok := req.Form["groupName"]; ok {
		groupName = req.Form["groupName"][0]
	}

	fmt.Println("id for delete = ", groupName)

	if err := db.RemoveGroup(groupName); err != nil {
		handleError(err, "Failed to remove group: %v\n", w)
		w.Write([]byte("OK: false"))
		return
	}
	w.Write([]byte("OK: true"))
}

func GetGroup(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var groupName string

	if _, ok := req.Form["groupName"]; ok {
		groupName = req.Form["groupName"][0]
	}

	rs, err := db.GetOneGroup(groupName)

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

func AddUserToGroup(w http.ResponseWriter, req *http.Request) { // add user (by ID) into group (by id)
	req.ParseForm()
	var groupName string
	var login string

	groupName = req.Form["groupName"][0]
	login = req.Form["login"][0]

	group, errG := db.GetOneGroup(groupName)
	user, errU := db.GetOneUser(login)
	if errG != nil {
		handleError(errG, "Failed to read database 'groups': %v", w)
		return
	}
	if errU != nil {
		handleError(errU, "Failed to read database 'users': %v", w)
		return
	}

	userExists := false
	for _, res := range group.Users {
		if res.Login == login {
			userExists = true
		}
	}

	if !userExists {
		if err := db.SaveUserIntoGroup(groupName, user); err != nil {
			handleError(err, "Failed to save group: %v into db groups", w)
			w.Write([]byte("OK: false"))
			return
		}
	}
	w.Write([]byte("OK: true"))
}

func GetUsersFromGroup(w http.ResponseWriter, req *http.Request) { // add user (by ID) into group (by id)
	req.ParseForm()
	groupName := req.Form["groupName"][0]

	group, err := db.GetOneGroup(groupName)

	if err != nil {
		handleError(err, "Failed to read database 'groups': %v", w)
		return
	}

	bs, err := json.Marshal(group.Users)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}
	w.Write(bs)
}

func DeleteUserFromGroup(w http.ResponseWriter, req *http.Request) { // delete user (by ID) from group (by id)
	req.ParseForm()
	groupName := req.Form["groupName"][0]
	login := req.Form["login"][0]

	user, _ := db.GetOneUser(login)

	if err := db.DeleteUserFromGroup(groupName, user); err != nil {
		handleError(err, "Failed to save group: %v into db groups", w)
		w.Write([]byte("OK: false"))
		return
	}
	w.Write([]byte("OK: true"))
}

func CopyUsersFromGroup(w http.ResponseWriter, req *http.Request) { // copy users from one group to another (by IDs)
	req.ParseForm()
	groupNameTo := req.Form["groupNameTo"][0]
	groupNameFrom := req.Form["groupNameFrom"][0]

	groupTo, err1 := db.GetOneGroup(groupNameTo)
	groupFrom, err2 := db.GetOneGroup(groupNameFrom)

	if err1 != nil {
		handleError(err1, "Failed to read database 'groups': %v", w)
		return
	}
	if err2 != nil {
		handleError(err2, "Failed to read database 'users': %v", w)
		return
	}

	isEmpty := len(groupTo.Users) == 0
	for _, resFrom := range groupFrom.Users {
		if isEmpty {
			user, _ := db.GetOneUser(resFrom.Login)
			if err := db.SaveUserIntoGroup(groupNameTo, user); err != nil {
				handleError(err, "Failed to save user: %v into db", w)
				w.Write([]byte("OK: false"))
				return
			}
		} else {
			hasUser := false
			for _, resTo := range groupTo.Users {
				if resTo.Login == resFrom.Login {
					hasUser = true
					break
				}
			}
			if !hasUser {
				user, _ := db.GetOneUser(resFrom.Login)
				if err := db.SaveUserIntoGroup(groupNameTo, user); err != nil {
					handleError(err, "Failed to save user: %v into db", w)
					w.Write([]byte("OK: false"))
					return
				}
			}
		}
	}
	w.Write([]byte("OK: true"))
}

func DeleteUsersFromGroup(w http.ResponseWriter, req *http.Request) { // delete users from first grope if users are in second group(by group id)
	req.ParseForm()
	groupNameForDel := req.Form["groupName1"][0]
	groupNameToCompare := req.Form["groupName2"][0]

	groupForDel, err1 := db.GetOneGroup(groupNameForDel)
	groupToCompare, err2 := db.GetOneGroup(groupNameToCompare)

	if err1 != nil {
		handleError(err1, "Failed to read database 'groups': %v", w)
		return
	}
	if err2 != nil {
		handleError(err2, "Failed to read database 'users': %v", w)
		return
	}

	var user db.User
	for _, userFrom := range groupToCompare.Users {
		for _, userToDel := range groupForDel.Users {
			if userFrom.Login == userToDel.Login {
				user, _ := db.GetOneUser(userToDel.Login)

				if err := db.DeleteUserFromGroup(groupNameForDel, user); err != nil {
					handleError(err, "Failed to delete user from group", w)
					w.Write([]byte("OK: false"))
					return
				}
				break
			}
		}
	}
	fmt.Println(user)
	w.Write([]byte("OK: true"))
}