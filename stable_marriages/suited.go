package main

type Suited struct {
	ID int
	Preferences []int
	HeldSuitor *Suitor
	NewProposals []*Suitor
}

func NewSuited(id int, preferences []int) (output Suited){
	output.ID = id
	output.Preferences = preferences
	return
}

func (s *Suited) Reject() {
	minRank := -1
	for _, suitor := range s.NewProposals {
		if s.Rank(suitor.ID) < minRank && s.Rank(suitor.ID) != -1{
			minRank = s.Rank(suitor.ID)
		}
	}
	for _, suitor := range s.NewProposals {
		if s.Rank(suitor.ID) == minRank {
			s.HeldSuitor = suitor
		} else {
			suitor.Reject()
		}
	}
}

func (s *Suited) Rank(ranked int) (output int){
	for _, value := range s.Preferences {
		if ranked == value {
			output = ranked
			return
		}
	}
	return -1
}


