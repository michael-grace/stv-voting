package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Vote symbolises an individual vote for a position with some number of places, and some candidates
// For URY elections, this won't automatically include RON (re-open nominations)
type Vote struct {
	Position     string   `yaml:"position"`
	NumPositions int      `yaml:"numPositions"`
	Candidates   []string `yaml:"candidates"`
}

// ElectionSet holds the data about an election set from the YAML file
type ElectionSet struct {
	ElectionSetName string `yaml:"electionSetName"`
	ControlPass     string `yaml:"controlPass"`
	VoterPass       string `yaml:"voterPass"`
	Votes           []Vote `yaml:"votes"`
}

func main() {
	// Get YAML Data
	data, err := ioutil.ReadFile("elections.yaml")
	if err != nil {
		panic(err)
	}

	var electionConfig []ElectionSet

	err = yaml.Unmarshal(data, &electionConfig)
	if err != nil {
		panic(err)
	}

	fmt.Println(electionConfig)

}
