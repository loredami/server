package auth

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func AddAuthRoutes(prefix string, router *mux.Router) {

	router.HandleFunc(prefix+"login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "login url!")
	}).Methods("POST")

	router.HandleFunc(prefix+"signup", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "signup url!")
	}).Methods("POST")

	router.HandleFunc(prefix+"activate", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "activate url!")
	}).Methods("GET")

	router.HandleFunc(prefix+"recover", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "recover url!")
	}).Methods("POST")
}
