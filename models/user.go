package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	PasswordEnc string    `json:"-"`
	Claims      []Claim   `json:"claims"`
}

type Claim struct {
	Key   string
	Value string
}
