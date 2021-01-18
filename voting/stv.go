package voting

var candidateDetails map[string]*candidate = make(map[string]*candidate)
var ballots []*ballot
var requiredPercentage float32

func redistributeVotes(from *candidate, elimination bool) {
	var voteCount int = 0
	for _, ballot := range ballots {
		lastVote := ballot.lastVote()
		if lastVote != nil && lastVote.candidate == from.name {
			voteCount++
		}
	}

	var percentPerVote float32
	if elimination {
		percentPerVote = float32(from.gainedPercent / float32(voteCount))
	} else {
		percentPerVote = float32((from.gainedPercent - requiredPercentage) / float32(voteCount))
	}

	for _, ballot := range ballots {
		lastVote := ballot.lastVote()
		if lastVote != nil && lastVote.candidate == from.name {
			b := ballot.nextVote()
			if b != nil {
				if !candidateDetails[b.candidate].won {
					candidateDetails[b.candidate].gainedPercent += percentPerVote
				}
				b.counted = true
			}
		}
	}
}

func runElection(numSeats int) (results []string) {

	// Count Winners
	var winners int = 0

	// First Preferences
	var percentPerVote float32 = float32(100 / len(ballots))

	for _, ballot := range ballots {
		b := ballot.nextVote()
		if b != nil {
			candidateDetails[b.candidate].gainedPercent += percentPerVote
			b.counted = true
		}
	}

	// Continuing Rounds
	for winners < numSeats {

		if len(candidateDetails) == numSeats {
			for _, nowWinner := range candidateDetails {
				nowWinner.won = true
			}
			break
		}

		var overThreshold []*candidate
		for _, candidate := range candidateDetails {
			if !candidate.won && candidate.gainedPercent >= requiredPercentage {
				candidate.won = true
				winners++
				overThreshold = append(overThreshold, candidate)
			}
		}

		if len(overThreshold) != 0 {
			// Votes over winning threshold
			for _, candidate := range overThreshold {
				redistributeVotes(candidate, false)
			}
		} else {
			// Eliminating Worst Candidate
			var worstCandidate *candidate = &candidate{gainedPercent: 101}
			for _, candidate := range candidateDetails {
				if candidate.gainedPercent < worstCandidate.gainedPercent {
					worstCandidate = candidate
				}
			}

			redistributeVotes(worstCandidate, true)

			delete(candidateDetails, worstCandidate.name)

			for _, b := range ballots {
				b.cleanse()
			}
		}
	}

	for name, candidate := range candidateDetails {
		if candidate.won {
			results = append(results, name)
		}
	}

	return

}

// STVElection runs the complexity of the STV election
func STVElection(numSeats int, candidates []string, pollCards [][]string) []string {

	// Required Percentage to Win
	requiredPercentage = float32(100 / numSeats)

	candidateDetails = make(map[string]*candidate)

	// Create Map
	for _, name := range candidates {
		candidateDetails[name] = &candidate{
			gainedPercent: 0,
			name:          name,
			won:           false,
		}
	}

	ballots = []*ballot{}

	for _, b := range pollCards {
		if len(b) > 0 {
			ballot := ballot{}
			for _, v := range b {
				ballot.votes = append(ballot.votes, &vote{candidate: v, counted: false})
			}
			ballots = append(ballots, &ballot)
		}
	}

	return runElection(numSeats)

}
