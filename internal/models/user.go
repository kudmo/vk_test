package models

// User
// Describes service user
type User struct {
	Id       int    `json:"-" gorm:"primary_key"`
	Login    string `json:"login" gorm:"unique; not null" `
	Password string `json:"-"`
}
