package entity

type User struct {
	Email       string `json:"email"`
	IsAnonymous bool   `json:"isAnonymous"`
	Phone       string `json:"phone"`
	UID         string `json:"uid"`
}
