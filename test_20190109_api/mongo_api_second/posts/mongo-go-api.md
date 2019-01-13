# MongoDB API with Go Lang

I've been working with MongoDB database for quite a long time now,
 and it proved effective with Node.js. But how exactly it measures up with Go?
 It actually is also very easy to do,
 and knowing how easy it is to deploy applications in Go,
 it's worth taking a look. So let's!

### Requirements

To create an application that will be working with MongoDB,
it obviously needs to use a specific database driver.
In Go I'm using `mgo` package.
And that might be all, because an `http` package is comes with Go by default.
But just like with Node's `express`,
there are tools that make those defaults easier to manage.
So that we use another library, `gorilla/mux` for easier creation of specific REST handlers.

For example, if you want to create two handlers under the same URL,
but different HTTP methods, you can do it using default package:

	http.HandleFunc("/api/items", func(w http.ResponseWriter, req *http.Request){
		if req.Method == "GET" {
			HandleFuncGET(w, req)
			return
		} else if req.Method == "POST" {
			HandleFuncPOST(w, req)
			return
		} else {
			HandleError(w, req)
		}
	})
	http.ListenAndServe(":3000", nil)

or a 3rd party alternative:

	r := mux.NewRouter()
	r.HandleFunc("/api/items", HandleFuncGET).Methods("GET")
	r.HandleFunc("/api/items", HandleFuncPOST).Methods("POST")
	http.ListenAndServe(":3000", r)
    
You can see the difference.

### Our REST API

In this example we want to create four simple REST actions: getting all items (GET),
getting one item (GET), saving an item (POST) and removing an item (DELETE).
We will follow a classic REST URL convention, as you can see below.

### MongoDB driver

Using the `mgo` driver is pretty easy,
but you need to remember that the methods that you are using, unless `mongoose` for Node,
are slightly different than Mongo's API.
The second thing you need to know is how to build a dynamic object representing eg.
Mongo query in a statically typed language as Go.
This is where `mgo/bson` package comes in place, which allows you to do that:

	db.C("items").Find(bson.M{"_id": id}).One(&res)

With MongoDB and Go there is one tricky part - an attribute `_id`.
You must remember to have a BSON mapping for our struct,
because _id is important in querying database using that BSON structure.
Our database struct looks like this then:

    type Item struct {
        ID    string `json:"id" bson:"_id,omitempty"`
        Value int    `json:"value"`
    }

### Separation of concerns

To have our application structured correctly,
we should divide our code into three files: an entry point with HTTP server (`main.go`),
all request handlers (`api.go`) and actions on database (`db.go`).

Our entry point looks just like in the description of `gorilla/mux` mentioned above:

    func main() {
        r := mux.NewRouter()
        r.HandleFunc("/api/items", api.GetAllItems).Methods("GET")
        r.HandleFunc("/api/items/{id}", api.GetItem).Methods("GET")
        r.HandleFunc("/api/items", api.PostItem).Methods("POST")
        r.HandleFunc("/api/items/{id}", api.DeleteItem).Methods("DELETE")

        http.ListenAndServe(":3000", r)
    }

Then we go to request handlers, and need to implement all functions.
Those functions have three things to do: call a database module,
handle its errors and write the result into `http.ResponseWriter` parameter.
The error handling takes most of the place in code, so it's trimmed a bit,
but it roughly looks like this:

    func GetAllItems(w http.ResponseWriter, req *http.Request) {
        rs, err := db.GetAll()
        bs, err := json.Marshal(rs)
        w.Write(bs)
    }

    func GetItem(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        id := vars["id"]
        rs, err := db.GetOne(id)
        bs, err := json.Marshal(rs)
        w.Write(bs)
    }

    func PostItem(w http.ResponseWriter, req *http.Request) {
        ID := req.FormValue("id")
        valueStr := req.FormValue("value")
        value, err := strconv.Atoi(valueStr)
        item := db.Item{ID: ID, Value: value}
        err = db.Save(item);
        w.Write([]byte("OK"))
    }

    func DeleteItem(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        id := vars["id"]
        err := db.Remove(id)
        w.Write([]byte("OK"))
    }

Then finally we can implement a communication with a database
(again, no error handling for clarity):

    var db *mgo.Database
    func init() {
        session, err := mgo.Dial("localhost/api_db")
        db = session.DB("api_db")
    }

    func collection() *mgo.Collection {
        return db.C("items")
    }

    func GetAll() ([]Item, error) {
        res := []Item{}
        if err := collection().Find(nil).All(&res); err != nil {
            return nil, err
        }
        return res, nil
    }

    func GetOne(id string) (*Item, error) {
        res := Item{}
        if err := collection().Find(bson.M{"_id": id}).One(&res); err != nil {
            return nil, err
        }
        return &res, nil
    }

    func Save(item Item) error {
        return collection().Insert(item)
    }

    func Remove(id string) error {
        return collection().Remove(bson.M{"_id": id})
    }

### Final result

The application is very simple both to implement and
to grasp (I hope at least) and you can easily extend it with other database models.
The final result weight about 8.1 megabytes.
Code is available [on Github](https://github.com/mycodesmells/mongo-go-api).
