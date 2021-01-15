package voting

type candidate struct {
	gainedPercent float32
	voters        []ballot
}

type ballot struct {
	id    int
	votes []string
}

// STVVoting takes the details for an election, and returns the elected candidates
func STVVoting(numSeats int, candidates []string, votes []ballot) (results []string) {
	var requiredPercent float32 = float32(100 / numSeats)

	var candidateDetails map[string]*candidate

	for _, name := range candidates {
		candidateDetails[name] = &candidate{
			gainedPercent: 0,
		}
	}

	for len(results) != numSeats {

		var percentPerVote float32 = float32(100 / len(votes))

		for i, b := range votes {
			var remains ballot = ballot{id: b.id, votes: b.votes[1:]}

			candidateDetails[b.votes[0]].voters = append(candidateDetails[b.votes[0]].voters, remains)

			votes[i].votes = votes[i].votes[1:]
			if len(votes[i].votes) == 0 {
				votes = append(votes[:i], votes[i+1:]...)
			}

			candidateDetails[b.votes[0]].gainedPercent += percentPerVote

		}

		var runoff bool = false

		for name, details := range candidateDetails {
			if details.gainedPercent >= requiredPercent {
				results = append(results, name)
			}

			runoff = true

			// Excess Votes
			var distributionPercentage float32 = details.gainedPercent - requiredPercent
			var reallocatePerVote float32 = distributionPercentage / float32(len(details.voters))

			for i, b := range details.voters {

				_, ok := candidateDetails[b.votes[0]]
				for !ok {
					b.votes = b.votes[1:]
					if len(b.votes) == 0 {
						details.voters = append(details.voters[:i], details.voters[i+1:]...)
					} else {
						_, ok = candidateDetails[b.votes[0]]
					}
				}

				candidateDetails[b.votes[0]].gainedPercent += reallocatePerVote

				var remains ballot = ballot{id: b.id, votes: b.votes[1:]}
				candidateDetails[b.votes[0]].voters = append(candidateDetails[b.votes[0]].voters, remains)

				// TODO Continue
				// Also, work out where I've actually got to
				// This is starting to get very ugly

			}

			// Yeet from map
		}

		if !runoff {
			// TODO

			// Yeet the worst
		}

	}

	return

}
