package main

import "log"

func main() {
	suitors := []Suitor{}
	suited := []Suited{}

	if len(suitors) != len(suited) {
		log.Fatal("Not equal - someone will be all alone!")
	}
	outerloop:
	for  {
		for _, suited := range suited {
			suited.NewProposals = []*Suitor{}
		}
		for _, suitor := range suitors {
			suited[suitor.Preferences[suitor.NextProposal]].NewProposals = append(suited[suitor.Preferences[suitor.NextProposal]].NewProposals, &suitor)
		}
		for _, suited := range suited {
			suited.Reject()
		}
		for _, suitor := range suitors {
			if suitor.NextProposal >= len(suitors) {
				break outerloop
			}
		}
	}
}
