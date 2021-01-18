package web

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
)

// Vote symbolises an individual vote for a position with some number of places, and some candidates
// For URY elections, this won't automatically include RON (re-open nominations)
type Vote struct {
	id           string
	Position     string   `yaml:"position"`
	NumPositions int      `yaml:"numPositions"`
	Candidates   []string `yaml:"candidates"`
	open         bool
	ballots      [][]string
}

// ElectionSet holds the data about an election set from the YAML file
type ElectionSet struct {
	id              string
	ElectionSetName string  `yaml:"electionSetName"`
	ControlPass     string  `yaml:"controlPass"`
	VoterPass       string  `yaml:"voterPass"`
	Votes           []*Vote `yaml:"votes"`
}

// ElectionConfig holds all the data from the YAML file about the elections
var ElectionConfig []*ElectionSet

// PopulateElections reads the YAML file, and fills in details
func PopulateElections() {
	// Get YAML Data
	data, err := ioutil.ReadFile("elections.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &ElectionConfig)
	if err != nil {
		panic(err)
	}

	// This is bad
	for _, election := range ElectionConfig {
		election.id = fmt.Sprintf("%x", rand.Intn(16777216))
		for _, vote := range election.Votes {
			vote.id = fmt.Sprintf("%x", rand.Intn(16777216))
		}
	}

	// Remove this later on
	fmt.Println(ElectionConfig[0].id)

}
