package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"./cassandra"
	"./globals"
	"./models"
	"./utils"
	"github.com/gorilla/mux"
)

var gS = globals.GlobalStates{}

// our main function
func main() {

	fmt.Println("app started")
	cassandra.Init(os.Args[1])
	router := mux.NewRouter()
	router.HandleFunc("/crudquery", CRUDQuery).Methods("POST")
	router.HandleFunc("/executequery", ExecuteQuery).Methods("POST")
	router.HandleFunc("/getstatus", GetStatus).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func CRUDQuery(w http.ResponseWriter, r *http.Request) {
	gS.IncrementRQ()

	session := cassandra.Session
	var data models.ReqCRUD
	utils.DecodeMessage(r, &data)
	responseType := r.Header.Get("responseType")

	data.Operation = strings.ToLower(data.Operation)
	query := utils.CreateCrudQueryString(data)
	var response models.Response

	if data.Operation == "select" {
		iter := session.Query(query).Iter()

		var row string
		var resultUnmarshaled []interface{}

		for iter.Scan(&row) {
			rowUnmarshaled := make(map[string]interface{})
			json.Unmarshal([]byte(row), &rowUnmarshaled)
			resultUnmarshaled = append(resultUnmarshaled, rowUnmarshaled)
		}
		errMessage := iter.Close()

		if errMessage != nil {
			response.ErrorMessage = errMessage.Error()
			w.WriteHeader(http.StatusBadRequest)
		}
		byteresp, _ := utils.EncodeResponse(resultUnmarshaled, responseType)
		response.Result = string(byteresp)

		res, _ := utils.EncodeResponse(&response, responseType)
		w.Write(res)
	} else {
		if data.Operation == "insert" || data.Operation == "update" || data.Operation == "delete" {
			err := session.Query(query).Exec()
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				response.ErrorMessage = err.Error()
				gS.IncrementFQ()
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Wrong type of operation.")
			gS.IncrementFQ()
		}
	}

	gS.DecrementRQ()
	gS.IncrementTQ()
	gS.IncrementSQ()

	// fmt.Fprintln(w, "Welcome!")
}

func ExecuteQuery(w http.ResponseWriter, r *http.Request) {
	gS.IncrementRQ()
	session := cassandra.Session
	var data models.ExecuteQuery
	var response models.Response

	utils.DecodeMessage(r, &data)
	err := session.Query(data.Query).Exec()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response.ErrorMessage = err.Error()
		gS.IncrementFQ()
	}
	gS.DecrementRQ()
	gS.IncrementTQ()
	gS.IncrementSQ()
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println(gS)
	keys, ok := r.URL.Query()["responseType"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := utils.EncodeResponse(&gS, keys[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(response)
}
