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
	router.HandleFunc("/controller/{id}/{voteid}/open", web.OpenVote)
	// router.HandleFunc("/controller/{id}/{voteid}/close", web.CloseVote)

	router.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) { web.Serve404(w) })
	http.Handle("/", router)
	http.ListenAndServe(":3000", router)

}
