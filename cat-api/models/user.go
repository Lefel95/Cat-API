package models

//UserLogin is the struct responsible for login credentials
type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
