package models

type SimpleUser struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Claims   []SimpleUserClaim `json:"claims"`
}

type SimpleUserClaim struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
