package api

import (
	// "fmt"
	"encoding/json"
	"gitlab.com/ravi.chandran/cartoon/data"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	var id int
	// fmt.Println(strings.SplitN(r.URL.Path, "/", 5)[1])
	if parts := strings.Split(r.URL.Path, "/"); len(parts) == 5 {
		// fmt.Println(parts)
		id, _ = strconv.Atoi(parts[4]) // invalid ID handled below
	}
	if id == 0 {
		http.Error(w, "No valid ID present", 400)
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
	var id int
	// fmt.Println(strings.SplitN(r.URL.Path, "/", 5)[1])
	if parts := strings.Split(r.URL.Path, "/"); len(parts) == 5 {
		// fmt.Println(parts)
		id, _ = strconv.Atoi(parts[4]) // invalid ID handled below
	}
	if id == 0 {
		http.Error(w, "No valid ID present", 400)
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

//  Handler for query params NOT USED CURRENTLY
func customHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))

}
