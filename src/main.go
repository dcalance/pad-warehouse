package main

import (
	"./cassandra"
	"./models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// our main function
func main() {

	fmt.Println("app started")
	cassandra.Init()

	router := mux.NewRouter()
	router.HandleFunc("/executequery", ExecuteQuery).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func ExecuteQuery(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	session := cassandra.Session

	var data models.ReqQuery
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(data)
	switch data.Operation {
	case "select":
		// fmt.Println("select")
		iter := session.Query("SELECT json * FROM test;").Iter()

		var suka string
		iter.Scan(&suka)
		fmt.Println(suka)

	case "insert":
		fmt.Println("insert")
	case "update":
		fmt.Println("update")
	case "delete":
		fmt.Println("delete")
	}

	// w.WriteHeader(http.StatusBadRequest)
	// fmt.Fprintln(w, "Welcome!")
}
