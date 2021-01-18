package voting

type candidate struct {
	gainedPercent float32
	won           bool
	name          string
}

type vote struct {
	candidate string
	counted   bool
}

type ballot struct {
	votes []*vote
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
