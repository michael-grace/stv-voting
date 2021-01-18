package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// ControllerElectionList is the handler for listing to a controller the votes
func ControllerElectionList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		serve400(w)
		return
	}

	var currentElection *ElectionSet

	for _, election := range ElectionConfig {
		if election.id == id {
			currentElection = election
			break
		}
	}

	if currentElection == nil {
		w.WriteHeader(http.StatusNotFound)
		serve404(w)
		return
	}

	fmt.Fprintf(w, "%v", currentElection)

}
