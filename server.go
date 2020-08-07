package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gitlab.com/ravi.chandran/cartoon/handler"
	"log"
	"net/http"
)

func main() {
	var router = mux.NewRouter()
	var api = router.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	var api1 = api.PathPrefix("/v1").Subrouter()
	api1.HandleFunc("/cartoons", handler.Handler)
	api1.HandleFunc("/cartoons/{id}", handler.Handler)
	api1.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	fmt.Println("Serving from http://localhost:8080/api/v1")
	log.Fatal(http.ListenAndServe(":8080", router))
}
