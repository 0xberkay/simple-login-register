package models

import "time"

type User struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"unique"`
	Points     int       `json:"points"`
	Referance  string    `json:"referance"`
	Referances int       `json:"referances"`
	Code       string    `json:"-"`
	Verified   bool      `json:"verified"`
	Created    time.Time `json:"created"`
	Password   []byte    `json:"-"`
}

type Order struct {
	Id    uint   `json:"id"`
	Email string `json:"email" gorm:"unique"`
	Order string `json:"order"`
	Price string `json:"price"`
}
