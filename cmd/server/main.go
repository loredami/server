package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/loredami/server/pkg/auth"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello world!")
	}).Methods("GET")
	auth.AddAuthRoutes("/", router)

	log.Fatal(http.ListenAndServe(":80", router))
}
