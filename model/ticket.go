package model

type TicketList struct {
	ID          int    `json:"id" db:"id"`
	Tag         string `json:"tag"  db:"tag"`
	TagID       int    `json:"tag_id" db:"tag_id"`
	Topic       string `json:"topic" db:"topic"`
	Location    string `json:"location" db:"location"`
	Tutor       string `json:"tutor" db:"tutor"`
	Amount      string `json:"amount" db:"amount"`
	Full        int    `json:"full" db:"max"`
	Date        string `json:"date" db:"date"`
	StartTime   string `json:"start_time" db:"start_time"`
	StopTime    string `json:"stop_time" db:"stop_time"`
	Price       int    `json:"price" db:"price"`
	Ticket      int    `json:"ticket" db:"access_code"`
	Description string `json:"description" db:"description"`
}
