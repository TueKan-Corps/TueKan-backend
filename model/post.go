package model

type CreatePost struct {
	AccountID   int    `json:"account_id"`
	Topic       string `json:"topic" db:"topic"`
	Location    string `json:"location" db:"location"`
	Date        string `json:"date"`
	StartTime   string `json:"start_time" db:"start_at"`
	StopTime    string `json:"stop_time" db:"end_at"`
	Category    int    `json:"category" db:"tag_id"`
	Price       int    `json:"price" db:"price"`
	Max         int    `json:"max" db:"max_participant"`
	Description string `json:"description" db:"description"`
}

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

//Posting for your posting data
type Posting struct {
	ID          int    `json:"id" db:"id"`
	Tag         string `json:"tag " db:"tag"`
	TagID       int    `json:"tag_id" db:"tag_id"`
	Topic       string `json:"topic" db:"topic"`
	Location    string `json:"location" db:"location"`
	Tutor       string `json:"tutor" db:"tutor"`
	Amount      int    `json:"amount" db:"amount"`
	Date        string `json:"date"`
	Full        int    `json:"full" db:"max"`
	StartTime   string `json:"start_time" db:"start_at"`
	StopTime    string `json:"stop_time" db:"stop_at"`
	Price       int    `json:"price" db:"price"`
	Participant string `json:"participant" db:"list"`
	Description string `json:"description" db:"description"`
}
