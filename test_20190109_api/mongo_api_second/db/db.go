package db

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type User struct {
	Login string `json:"_id" bson:"_id"`
	Pass  string `json:"pass" bson:"pass"`
	Age   int    `json:"age" bson:"age"`
}

var DBusers, DBgroups *mgo.Database

func init() {
	fmt.Println("start func Init()")
	session, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DBusers, DBgroups = session.DB("users"), session.DB("groups")
}

func collectionUsers() *mgo.Collection {
	return DBusers.C("users")
}

func GetAllUsers() ([]User, error) { // returns all users from the database
	res := []User{}

	if err := collectionUsers().Find(nil).All(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func SaveUser(user User) error {  // Save a user to the DB users
	return collectionUsers().Insert(user)
}

func RemoveUser(id string) error {  // Remove a user by id(login) from DB users
	return collectionUsers().RemoveId(id)
}

func GetOneUser(id string) (*User, error) {  // get user by id from DB users
	res := User{}

	if err := collectionUsers().FindId(id).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func FindIdByLogin(id string) (*User, error) {  // get user's ID by login from DB users
	res := User{}

	if err := collectionUsers().FindId(id).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}




///////////////////


type Group struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Value int    `json:"value"`
}

func collectionGroups() *mgo.Collection {
	return DBgroups.C("groups")
}

func GetAllGroups() ([]Group, error) {
	res := []Group{}

	if err := collectionUsers().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func GetOneGroup(id string) (*Group, error) {
	res := Group{}

	if err := collectionGroups().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func SaveGroup(group Group) error {
	return collectionGroups().Insert(group)
}

func RemoveGroup(id string) error {
	return collectionGroups().Remove(bson.M{"_id": id})
}
