package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"./cassandra"
	"./models"
	"github.com/gorilla/mux"
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
	data.Operation = strings.ToLower(data.Operation)
	query := createQueryString(data)
	// fmt.Println(data)
	if data.Operation == "select" {
		iter := session.Query(query).Iter()

		var row string
		var resultUnmarshaled []map[string]interface{}

		for iter.Scan(&row) {
			rowUnmarshaled := make(map[string]interface{})
			json.Unmarshal([]byte(row), &rowUnmarshaled)
			resultUnmarshaled = append(resultUnmarshaled, rowUnmarshaled)
		}
		errMessage := iter.Close()
		var response models.Response

		if errMessage != nil {
			response.ErrorCode = 1
			response.ErrorMessage = errMessage.Error()
		} else {
			response.ErrorCode = 0
		}
		response.Result = resultUnmarshaled

		res, _ := json.Marshal(response)
		w.Write(res)
	} else {
		if data.Operation == "insert" || data.Operation == "update" || data.Operation == "delete" {
			err := session.Query(query).Exec()
			fmt.Println(err)
		} else {
			fmt.Println("Wrong type of operation.")
		}
	}

	// w.WriteHeader(http.StatusBadRequest)
	// fmt.Fprintln(w, "Welcome!")
}

func createQueryString(data models.ReqQuery) string {
	resString := data.Operation + strings.Join(data.Columns, ", ") + " FROM " + data.Table
	if data.Join != "" {
		resString += data.Join + " ON " + data.JoinCondition
	}
	if data.Conditions != "" {
		resString += " WHERE " + data.Conditions
	}
	if data.Orderby != "" {
		resString += " ORDER BY " + data.Orderby
	}
	if data.Limit != 0 {
		resString += " LIMIT "
	}
	if data.AllowFiltering {
		resString += " ALLOW FILTERING"
	}
	resString += ";"
	return resString
}
