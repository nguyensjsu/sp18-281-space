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
		Addr:     "localhost:6379",
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
	router.HandleFunc("/", getIndexHandler).Methods("GET")
	//router.HandleFunc("/", postIndexHandler).Methods("POST")
	router.HandleFunc("/customerOrder", getCustomerOrderHandler).Methods("GET")
	router.HandleFunc("/customerOrder", postCustomerOrderHandler).Methods("POST")
router.HandleFunc("/reservation", getReservationHandler).Methods("GET")
	router.HandleFunc("/reservation", postReservationHandler).Methods("POST")

	//router.HandleFunc("/", postHandler).Methods("POST")

	router.HandleFunc("/login", loginGetHandler).Methods("GET")
	router.HandleFunc("/login", loginPostHandler).Methods("POST")
	router.HandleFunc("/signup", signUpPostHandler).Methods("POST")

	fileSer := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileSer))

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func getCustomerOrderHandler(writer http.ResponseWriter, re *http.Request) {
	comments, err := client.LRange("customer", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(writer, "CustomerOrder.html", comments)
}



func postReservationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reservation := r.PostForm.Get("reservation")
	client.RPush("reservation", reservation)
	http.Redirect(w, r, "/reservation", 302)
}


func getReservationHandler(writer http.ResponseWriter, re *http.Request) {
	comments, err := client.LRange("reservation", 0, 6).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(writer, "reservation.html", comments)
}



func postCustomerOrderHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	customer := r.PostForm.Get("customer")
	client.RPush("customer", customer)
	http.Redirect(w, r, "/customerOrder", 302)
}



func getIndexHandler(writer http.ResponseWriter, re *http.Request) {
	comments, err := client.LRange("customer", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(writer, "index.html", comments)
}
func loginGetHandler(writer http.ResponseWriter, re *http.Request) {
	templates.ExecuteTemplate(writer, "login.html", nil)
}

func signUpPostHandler(writer http.ResponseWriter, re *http.Request) {
	templates.ExecuteTemplate(writer, "signup.html", nil)
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
