package main

import "log"

func main() {
	suitors := []Suitor{}
	suited := []Suited{}

	if len(suitors) != len(suited) {
		log.Fatal("Not equal - someone will be all alone!")
	}

	for i := 0; i < len(suitors); i++ {
		for _, suited := range suited {
			suited.NewProposals = []*Suitor{}
		}
		for _, suitor := range suitors {
			suited[suitor.Preferences[i]].NewProposals = append(suited[suitor.Preferences[i]].NewProposals, &suitor)
		}
		for _, suited := range suited {
			suited.Reject()
		}
	}
}
