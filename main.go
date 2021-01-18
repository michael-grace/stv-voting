package main

import (
	"github.com/gorilla/mux"
	"github.com/michael-grace/stv-voting/web"
	"net/http"
)

func main() {

	web.PopulateElections()

	router := mux.NewRouter()

	router.HandleFunc("/controller/{id}", web.ControllerElectionList)

	http.Handle("/", router)
	http.ListenAndServe(":3000", router)

}
