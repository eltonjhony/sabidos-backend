package model

type AccountModel struct {
	DefaultAvatarId int    `json:"defaultAvatarId"`
	Name            string `json:"name"`
	NickName        string `json:"nickname"`
	Email           string `json:"email"`
	IsAnonymous     bool   `json:"isAnonymous"`
	Phone           string `json:"phone"`
	Uid             string `json:"uid"`
}
