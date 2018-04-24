package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocql/gocql"
)

type heartbeatResponse struct {
	Status string `json: "status"`
	Code   int    `json:"code"`
}

func main() {
	fmt.Println("Start ...")
	// aRouter := mux.NewRouter().StrictSlash(true)
	// aRouter.HandleFunc("/", myheartbeat)
	// log.Fatal(http.ListenAndServe(":8080", aRouter))

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "counterdb"
	session, _ := cluster.CreateSession()

	fmt.Println("Create Keyspace successfully")

	// Insert employeee
	if err := session.Query("INSERT INTO customer(customerid, customer_first, customer_last, customer_email, customer_password) VALUES (1, 'Hoang', 'Nguyen', 'hoang1127@gmail.com', 'abc123')").Exec(); err != nil {
		log.Fatal(err)
	}
	var firstname string
	fmt.Println("Insert into Customer table")

	if err := session.Query("SELECT customer_first FROM customer ").Scan(&firstname); err != nil {
		log.Fatal(err)
	}

	fmt.Println(firstname)

	defer session.Close()
}

func myheartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 300})
}
