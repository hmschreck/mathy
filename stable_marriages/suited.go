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
	for _, suitor := range s.NewProposals {
		if s.Rank(suitor.ID) < s.Rank(s.HeldSuitor.ID) && s.Rank(suitor.ID) != -1{
			s.HeldSuitor = suitor
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


