package models

type Users struct {
	ID         int64  `json:"id"`
	Firstname  string `json:"first_name"`
	Lastname   string `json:"last_name"`
	Middlename string `json:"middle_name"`
	Username   string `json:"user_name"`
	Status     string `json:"status"`
	Phone      int64  `json:"phone"`
	Email      string `json:"email"`
	// PhoneVerifiedAt time.Time `json:"phone_verified_at"`
	// EmailVerifiedAt time.Time `json:"email_verified_at"`
	Birthday string `json:"birthday"`
	Gender   int    `json:"gender"`
	Address  string `json:"address"`
	Avatar   string `json:"avatar"`
}
