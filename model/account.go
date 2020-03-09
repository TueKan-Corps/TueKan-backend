package model

// Account model
type Account struct {
	ID         int    `json:"id" db:"id"`
	Username   string `json:"username" db:"username"`
	Password   string `json:"password" db:"password"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	CoinAmount int    `json:"coin_amount" db:"coin_amount"`
	ProfileIMGPath string `json:"profile_img_path,omitempty" db:"profile_img_path,omitempty"`
}
