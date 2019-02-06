package main

type Suitor struct {
	ID int
	Preferences []int
}

func NewSuitor(id int, preferences []int) (output Suitor) {
	output.ID = id
	output.Preferences = preferences
	return
}

