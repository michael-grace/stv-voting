package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// VoterElectionList is the handler for letting voters know what votes are open
func VoterElectionList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		serve400(w)
		return
	}

	var currentElection *ElectionSet

	for _, election := range ElectionConfig {
		if election.ID == id {
			currentElection = election
			break
		}
	}

	if currentElection == nil {
		Serve404(w)
		return
	}

	if err := serveTemplate(w, "availableVotes.html", currentElection, 200); err != nil {
		fmt.Println(err)
	}

}
