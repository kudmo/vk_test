package models

// User
// Describes service user
type User struct {
	Login    string `json:"login" gorm:"primary_key" `
	Password string `json:"-"`
}
