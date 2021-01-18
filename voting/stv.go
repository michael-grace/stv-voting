package voting

import "fmt"

type candidate struct {
	gainedPercent float32
	won           bool
	name          string
	eliminated    bool
	voters        []ballot
}

var candidateDetails map[string]*candidate = make(map[string]*candidate)
var requiredPercentage float32

type vote struct {
	candidate string
	counted   bool
}
type ballot struct {
	id    int
	votes []*vote
}

var ballots []*ballot

func cleanseAllBallots() {
	for _, b := range ballots {
		b.cleanse()
	}
}

func (b *ballot) cleanse() {
	for i := len(b.votes) - 1; i >= 0; i-- {
		_, ok := candidateDetails[b.votes[i].candidate]
		if !ok {
			b.votes = append(b.votes[:i], b.votes[i+1:]...)
		}
	}

}

func (b *ballot) nextVote() *vote {
	for _, vote := range b.votes {
		if !vote.counted {
			return vote
		}
	}
	return nil
}

func (b *ballot) lastVote() *vote {
	for i := 0; i < len(b.votes); i++ {
		if !b.votes[i].counted {
			return b.votes[i-1]
		}
	}
	return nil
}

func (b *ballot) stillIn() bool {
	fmt.Printf("Ballot: ")
	for _, v := range b.votes {
		fmt.Printf("Vote: %v, %v ", v.counted, v.candidate)
	}
	fmt.Printf("\n")
	if len(b.votes) > 0 {
		return !b.votes[len(b.votes)-1].counted
	} else {
		return false
	}
}

func voteRedistribution(firstLevel bool) (winners int) {

	// Next Stage - Redistribution or Elimination
	var needsRedistributing []*candidate
	for _, candidate := range candidateDetails {
		if !candidate.won && candidate.gainedPercent >= requiredPercentage {
			candidate.won = true
			fmt.Printf("Winner: %v\n", candidate.name)
			winners++
			needsRedistributing = append(needsRedistributing, candidate)
		}
	}

	if len(needsRedistributing) != 0 {
		// Redistributing Results
		for _, goodCandidate := range needsRedistributing {
			fmt.Printf("Redistributing %v\n", goodCandidate.name)
			var voteCount int = 0
			for _, ballot := range ballots {
				lastVote := ballot.lastVote()
				if lastVote != nil {
					if lastVote.candidate == goodCandidate.name {
						voteCount++
					}
				}
			}
			var percentPerVote float32 = float32((goodCandidate.gainedPercent - requiredPercentage) / float32(voteCount))

			for _, ballot := range ballots {
				lastVote := ballot.lastVote()
				if lastVote != nil {
					if lastVote.candidate == goodCandidate.name {
						b := ballot.nextVote()
						if b != nil {
							fmt.Printf("Redistributed Vote for %v\n", b.candidate)
							candidateDetails[b.candidate].gainedPercent += percentPerVote
							b.counted = true
						}
					}
				}
			}
		}

		winners += voteRedistribution(false)

	} else if firstLevel {
		// Eliminate the Lowest (only if no redistribution happened before)
		fmt.Println("We're going into eliminations")
		var worstCandidate *candidate = &candidate{gainedPercent: 100}
		for _, candidate := range candidateDetails {
			if candidate.gainedPercent < worstCandidate.gainedPercent {
				worstCandidate = candidate
			}
		}

		var voteCount int = 0
		for _, ballot := range ballots {
			lastVote := ballot.lastVote()
			if lastVote != nil {
				if ballot.lastVote().candidate == worstCandidate.name {
					voteCount++
				}
			}
		}
		var percentPerVote float32 = float32(worstCandidate.gainedPercent / float32(voteCount))

		for _, ballot := range ballots {
			b := ballot.nextVote()
			if b != nil {
				candidateDetails[b.candidate].gainedPercent += percentPerVote
				b.counted = true
			}
		}

		delete(candidateDetails, worstCandidate.name)

		cleanseAllBallots()

		winners += voteRedistribution(false)

	}

	return
}

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

	// for winners < numSeats {

	// 	if len(candidateDetails) == numSeats {
	// 		for _, nowWinner := range candidateDetails {
	// 			nowWinner.won = true
	// 		}
	// 		break
	// 	}

	// 	fmt.Println("New Round")
	// 	fmt.Printf("In ths Round: %v\n", candidateDetails)

	// 	// Calculate Percentage Per Vote
	// 	var voteCount int
	// 	for _, val := range ballots {
	// 		if val.stillIn() {
	// 			voteCount++
	// 		}
	// 	}

	// 	fmt.Printf("Votes In: %v\n", voteCount)

	// 	if voteCount == 0 {
	// 		// End of Election
	// 		// fmt.Println("Ending the Election")
	// 		// for _, nowWinner := range candidateDetails {
	// 		// 	nowWinner.won = true
	// 		// }
	// 		// break
	// 		voteCount = 100
	// 	}

	var percentPerVote float32 = float32(100 / len(ballots))

	// First Votes
	for _, ballot := range ballots {
		b := ballot.nextVote()
		if b != nil {
			candidateDetails[b.candidate].gainedPercent += percentPerVote
			b.counted = true
		}
	}

	// Continuing Rounds
	for winners < numSeats {

		fmt.Printf("New Loop: In Race: %v\n", candidateDetails)

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

	// winners += voteRedistribution(true)

	// }

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
			eliminated:    false,
		}
	}

	ballots = []*ballot{}

	for _, b := range pollCards {
		ballot := ballot{}
		for _, v := range b {
			ballot.votes = append(ballot.votes, &vote{candidate: v, counted: false})
		}
		ballots = append(ballots, &ballot)
	}

	return runElection(numSeats)

}
