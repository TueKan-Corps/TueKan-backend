package model

//PostList end point return all post list
type PostList struct {
	ID        int    `json:"id" db:"id"`
	Tag       string `json:"tag " db:"tag"`
	TagID     int    `json:"tag_id" db:"tag_id"`
	Topic     string `json:"topic" db:"topic"`
	Location  string `json:"location" db:"location"`
	Tutor     string `json:"tutor" db:"tutor"`
	Amount    int    `json:"amount" db:"amount"`
	Date      string `json:"date"`
	Full      int    `json:"full" db:"max_participant"`
	StartTime string `json:"start_time" db:"start_at"`
	StopTime  string `json:"stop_time" db:"stop_at"`
	Price     int    `json:"price" db:"price"`
}
