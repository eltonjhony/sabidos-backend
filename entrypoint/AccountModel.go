package entrypoint

// "defaultAvatarId":10,
// "name":"Elton",
// "nickname":"eljholiveidradaaaao",
// "email":"",
// "isAnonymous":true,
// "phone":"",
// "uid":"yiXtigKxtEVKl5mB4q7ZKx"

type AccountModel struct {
	DefaultAvatarId int    `json:"defaultAvatarId"`
	Name            string `json:"name"`
	NickName        string `json:"nickname"`
	Email           string `json:"email"`
	IsAnonymous     bool   `json:"isAnonymous"`
	Phone           string `json:"phone"`
	Uid             string `json:"uid"`
}
