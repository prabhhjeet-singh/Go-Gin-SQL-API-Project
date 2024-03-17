package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func Save(e Event) {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
