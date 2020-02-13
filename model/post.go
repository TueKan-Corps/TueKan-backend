package model

//Post model
type Post struct {
	ID          int    `json:"id" db:"id"`
	AccountID   int    `json:"account_id" db:"account_id"`
	Topic       string `json:"topic" db:"topic"`
	Location    string `json:"location" db:"location"`
	Description string `json:"description" db:"description"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
	CreatedAt   string `json:"created_at" db:"created_at"`
}

//CreatePost model
type CreatePost struct {
	AccountID   int    `json:"account_id" from:"account_id" query:"account_id"`
	Topic       string `json:"topic" from:"topic" query:"topic"`
	Location    string `json:"location" from:"location" query:"location"`
	Description string `json:"description" from:"description" query:"description"`
	UpdatedAt   string `json:"updated_at" from:"updated_at" query:"updated_at"`
	CreatedAt   string `json:"created_at" from:"created_at" query:"created_at"`
}
