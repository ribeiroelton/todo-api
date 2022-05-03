package domain

import "time"

type ToDo struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Responsible string    `json:"responsible"`
	When        time.Time `json:"when"`
}
