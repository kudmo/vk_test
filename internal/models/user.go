package models

// User
// Describes Dozen user: Student or Tutor
type User struct {
	Login    string `json:"login" gorm:"primary_key" `
	Password string `json:"password"`
}
