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
		if election.ID == id {
			currentElection = election
			break
		}
	}

	if currentElection == nil {
		Serve404(w)
		return
	}

	if err := serveTemplate(w, "listElections.html", currentElection, 200); err != nil {
		fmt.Println(err)
	}

}

// OpenVote marks a vote as open and accepting responses
func OpenVote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		serve400(w)
		return
	}

	voteID, ok := vars["voteid"]
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

	for _, vote := range currentElection.Votes {
		if vote.ID == voteID {
			vote.Open = true
			http.Redirect(w, r, fmt.Sprintf("/controller/%v", id), 303)
		}
	}

	Serve404(w)

}
