package model

import "time"

type ToDo struct {
	Id          int
	Title       string
	Description string
	Responsible string
	When        time.Time
}
