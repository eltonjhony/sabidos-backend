package model

type UpdateAccountModel struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsAnonymous bool   `json:"isAnonymous"`
	Phone       string `json:"phone"`
}
