package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)

type Group struct {
	GroupName string `json:"_id" bson:"_id"`
	Users     []User `json:"users" bson:"users"`
}

var DBgroups *mgo.Database

func init() {
	fmt.Println("start func Init()")
	session, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DBgroups = session.DB("groups")
}

func collectionGroups() *mgo.Collection {
	return DBgroups.C("groups")
}

func GetAllGroups() ([]Group, error) {
	res := []Group{}

	if err := collectionGroups().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func SaveGroup(group Group) error {
	return collectionGroups().Insert(group)
}


func RemoveGroup(id string) error {
	return collectionGroups().Remove(bson.M{"_id": id})
}

func GetOneGroup(id string) (*Group, error) {
	res := Group{}

	if err := collectionGroups().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func SaveUserIntoGroup(id string, user *User) error {
	res := Group{}
	err := collectionGroups().Find(bson.M{"_id": id}).One(&res)
	if err != nil {
		log.Fatal(err)
	}

	update := bson.M{"$push": bson.M{"users": user}}
	err = collectionGroups().UpdateId(id,update)
	return err
}