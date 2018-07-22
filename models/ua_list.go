package models

type UserAdvicesList struct {
	ID     int         `json:"id"`
	OpenID string      `json:"openid"`
	List   interface{} `json:"list"`
}
