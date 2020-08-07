package main

import (
	"fmt"
	"gitlab.com/ravi.chandran/cartoon/api"
	"log"
	"net/http"
)

func generalHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[8:])
	switch r.URL.Path[8:] {
	case "cartoons":
		api.Handler(w, r)
	case "":
		w.Write([]byte(`{"message": "Invalid Resource"}`))
	default:
		w.Write([]byte(`{"message": "Invalid Resource"}`))
	}
}

func main() {
	http.HandleFunc("/api/v1/", generalHandler)
	fmt.Println("Serving from http://localhost:8080/api/v1")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
