package main

type Suitor struct {
	ID int
	Preferences []int
	NextProposal int
}

func NewSuitor(id int, preferences []int) (output Suitor) {
	output.ID = id
	output.Preferences = preferences
	return
}

func (s *Suitor) Reject() {

}

