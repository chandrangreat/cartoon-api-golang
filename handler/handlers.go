package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/ravi.chandran/cartoon/data"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var cartoonData = data.LoadData()

// Handler for taking requests
func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	case "PUT":
		putHandler(w, r)
	case "DELETE":
		deleteHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(cartoonData)

	if err != nil {
		log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println("error writing result", err)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var cartoon data.Cartoon
	json.Unmarshal(reqBody, &cartoon)
	cartoonData = append(cartoonData, cartoon)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cartoon)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID format")
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedCartoon data.Cartoon
	json.Unmarshal(reqBody, &updatedCartoon)
	for index, cartoon := range cartoonData {
		if cartoon.ID == id {
			cartoonData[index] = updatedCartoon
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCartoon)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println("Invalid ID format")
		return
	}
	for index, cartoon := range cartoonData {
		if cartoon.ID == id {
			cartoonData = removeIndex(cartoonData, index)
		}
	}
	w.WriteHeader(http.StatusOK)
}

func removeIndex(s []data.Cartoon, index int) []data.Cartoon {
	return append(s[:index], s[index+1:]...)
}
