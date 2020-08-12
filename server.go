package main

import (
	"fmt"
	"gitlab.com/ravi.chandran/cartoon/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/cartoons/", api.Handler)
	fmt.Println("Serving from http://localhost:8080/api/v1")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
