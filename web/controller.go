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
		Serve404(w)
		return
	}

	if err := serveTemplate(w, "test.html", nil, 200); err != nil {
		fmt.Println(err)
	}

}
