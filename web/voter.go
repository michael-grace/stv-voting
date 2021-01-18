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

// VoterChoice is the handler for a vote selection page
func VoterChoice(w http.ResponseWriter, r *http.Request) {
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

	// Urghh...this is a mess
	for _, vote := range currentElection.Votes {
		if vote.ID == voteID {
			if vote.Open {
				if err := serveTemplate(w, "voterChoice.html", struct {
					Vote *Vote
					Nums []string
				}{
					Vote: vote,
					Nums: generateVoteList(len(vote.Candidates)),
				}, 200); err != nil {
					fmt.Println(err)
				}
				break
			} else {
				serve403(w)
			}
		}
	}
}
