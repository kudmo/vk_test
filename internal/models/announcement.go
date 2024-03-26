package models

import "time"

// Announcement
//
// Describes an ad on the marketplace
type Announcement struct {
	Id        int       `json:"-" gorm:"primary_key"`
	ClientId  int       `json:"-" gorm:"not null"`
	Client    User      `json:"-" gorm:"not null"`
	Title     string    `json:"title" gorm:"not null;type:varchar(50);default:null"`
	Text      string    `json:"text" gorm:"not null;type:varchar(256);default:null"`
	ImageRef  string    `json:"image_ref" gorm:"not null"`
	Price     int       `json:"price" gorm:"not null"`
	CreatedAt time.Time `json:"-"`
}
