package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	session *mgo.Session
	err     error
	dbhost  = "localhost"
)

type (
	PersonStruct = struct {
		ID           string `json:"_id" bson:"_id"`
		Name         string `json:"name" bson:"name"`
		Email        string `json:"email" bson:"email"`
		MobileNumber string `json:"mobile_number" bson:"mobile_number"`
	}
	obj map[string]interface{}
)

func init() {

	session, err = mgo.Dial(dbhost)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	go func() {
		for range time.Tick(time.Second) {
			сonnectionCheck()
		}
	}()
}

func main() {
	fmt.Println("TEST")

	//user := PersonStruct{"11", "Petghesadsar", "emaasdasdasilP", "06733333"}
	//AddUser(user)

	//user := PersonStruct{"1", "PeterSON", "emailPNew", "11111111"}
	//UpdateUser(user)

	//RemoveUserById("10")

	//person, _ := GetUserByIndex("2")
	//fmt.Printf("person => %+v", person)

	exists, err := CheckByIndex("10")
	fmt.Printf("Person exists? => %t.  Error => %v", exists, err)
}

// DB wrapper for mgo.Session.DB
func DB(dname string) *mgo.Database {
	сonnectionCheck()
	return session.DB(dname)
}

// сonnectionCheck reconect on lose conect
func сonnectionCheck() {
	if err := session.Ping(); err != nil {
		fmt.Println("Lost connection to db!")
		session.Refresh()
		if err := session.Ping(); err == nil {
			fmt.Println("Reconnect to db successful.")
		}
	}
}

func AddUser(person PersonStruct) {
	err := DB("test").C("testUser").Insert(person)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateUser(person PersonStruct) {
	//err :=  DB("test").C("testUser").Update(obj{"_id": person.ID}, person) // обновляет все структуру по объкту, а не одно поле
	//err :=  DB("test").C("testUser").UpdateId(person.ID, person) // обновляет все структуру по ID, а не одно поле
	err := DB("test").C("testUser").UpdateId(person.ID, obj{"$set": obj{"name": person.Name, "email": person.Email, "mobile_number": person.MobileNumber}})
	if err != nil {
		fmt.Println(err)
	}
}

func RemoveUserById(id string) {
	err := DB("test").C("testUser").RemoveId(id)
	if err != nil {
		fmt.Println(err)
	}
}

func GetUserByIndex(id string) (*PersonStruct, error) {
	result := PersonStruct{}

	err := DB("test").C("testUser").FindId(id).One(&result)
	if err != nil {
		fmt.Println(err)
	}

	return &result, nil
}

func CheckByIndex(id string) (bool, error) {
	result := PersonStruct{}
	err := DB("test").C("testUser").FindId(id).One(&result)
	if err != nil {
		return false, err
	}
	return true, err
}
