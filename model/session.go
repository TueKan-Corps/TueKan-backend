package model

// Key model
type Session struct {
	ID        int    `json:"id"`
	AccountID int    `json:"account_id"`
	Token     string `json:"token"`
}
