package model

//Post model
type Post struct {
	ID             int    `json:"id" db:"id"`
	AccountID      int    `json:"account_id" db:"account_id"`
	Topic          string `json:"topic" db:"topic"`
	Location       string `json:"location" db:"location"`
	Description    string `json:"description" db:"description"`
	UpdatedAt      string `json:"updated_at" db:"updated_at"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	HeldAt         string `json:"held_at" db:"held_at"`
	Tag            string `json:"tag" db:"tag"`
	MaxParticipant int    `json:"max_participant" db:"max_participant"`
}
