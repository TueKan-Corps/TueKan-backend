package model

// Account model
type Account struct {
	ID         int    `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	Password   string `json:"password" db:"password"`
	CoinAmount int    `json:"coin_amount" db:"coin_amount"`
}
