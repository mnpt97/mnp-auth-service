package models

type User struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	PasswordEnc string  `json:"-"`
	Claims      []Claim `json:"claims"`
}

type Claim struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
