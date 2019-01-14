package db

import (
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Login string    `json:"_id" bson:"_id"`
	Pass string    `json:"pass" bson:"pass"`
	Age int    `json:"age" bson:"age"`
}

var DBusers, DBgroups *mgo.Database

func init() {
	session, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DBusers, DBgroups = session.DB("users"), session.DB("groups")
}

// GetAll returns all items from the database
func GetAllUsers() ([]User, error) {
	res := []User{}

	if err := collectionUsers().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// Save an item to the database
func SaveUser(user User) error {
	return collectionUsers().Insert(user)
}

// Remove an item from the database
func RemoveUser(id string) error {
	return collectionUsers().RemoveId(id)
}

// GetOne returns a single item from the database.
func GetOneUser(id string) (*User, error) {
	//res2 := collectionUsers().FindId(id)
	//fmt.Println("res ", res2)
	//fmt.Printf("type of res = %T ", res2)
	res := User{}

	if err := collectionUsers().Find(id).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

//TODO +++

type Group struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Value int    `json:"value"`
}


func collectionUsers() *mgo.Collection {
	return DBusers.C("users")
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
