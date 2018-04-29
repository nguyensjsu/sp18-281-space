package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var templates *template.Template
var client *redis.Client

func main() {
	fmt.Println("Start ...")

	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	templates = template.Must(template.ParseGlob("templates/*.html"))
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(writer http.ResponseWriter, re *http.Request) {
	comments, err := client.LRange("customer", 0, 10).Result()
	if err != nil {
		return
	}
	templates.ExecuteTemplate(writer, "index.html", comments)
}
