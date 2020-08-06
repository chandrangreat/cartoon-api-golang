package api

import (
	// "fmt"
	"encoding/json"
	"gitlab.com/ravi.chandran/cartoon/data"
	"io/ioutil"
	"log"
	"net/http"
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
