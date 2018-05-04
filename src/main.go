package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var templates *template.Template
var client *redis.Client
var store = sessions.NewCookieStore([]byte("password1"))

func main() {
	fmt.Println("Start ...")

	//	client = redis.NewClient(&redis.Options{
	//		Addr: "127.0.0.1:6379",
	//	})

	client = redis.NewClient(&redis.Options{
		Addr:     "54.67.84.81:6379",
		Password: "",
		DB:       0, // use default DB
	})

	/*	func NewRedisServer() *redis.Client {
			client := redis.NewClient(&redis.Options{
				Addr:     "13.57.96.105:6379",
				Password: "foobared",
				DB:       0,  // use default DB
			})

		  return client
		}
	*/

	templates = template.Must(template.ParseGlob("templates/*.html"))
	router := mux.NewRouter()
	router.HandleFunc("/", getHandler).Methods("GET")
	router.HandleFunc("/", postHandler).Methods("POST")

	router.HandleFunc("/login", loginGetHandler).Methods("GET")
	router.HandleFunc("/login", loginPostHandler).Methods("POST")
	router.HandleFunc("/test", testGetHandler).Methods("GET")

	fileSer := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileSer))

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func getHandler(writer http.ResponseWriter, re *http.Request) {
	comments, err := client.LRange("customer", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(writer, "index.html", comments)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	customer := r.PostForm.Get("customer")
	client.LPush("customer", customer)
	http.Redirect(w, r, "/", 302)
}

func loginGetHandler(writer http.ResponseWriter, re *http.Request) {
	templates.ExecuteTemplate(writer, "login.html", nil)
}

func loginPostHandler(writer http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.PostForm.Get("usernname")
	session, _ := store.Get(req, "session")

	session.Values["username"] = username
	session.Save(req, writer)
}

func testGetHandler(writer http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")
	untyped, ok := session.Values["username"]

	if !ok {
		return
	}

	username, ok := untyped.(string)
	if !ok {
		return
	}
	writer.Write([]byte(username))
}
