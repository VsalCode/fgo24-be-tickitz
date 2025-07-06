package dto;

import (
	"time"
)

type Ticket struct {
	Title  string    `json:"title"`
	Genre  []string  `json:"genre"`
	Date   time.Time `json:"date"`
	Time   time.Time `json:"time"`
	Seat   []string  `json:"seat"`
	Cinema string    `json:"cinema"`
}
