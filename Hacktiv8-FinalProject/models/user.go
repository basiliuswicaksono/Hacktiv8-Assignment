package models

import (
	"finalProject/helpers"

	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" valid:"required~Email is required, email~Invalid format email"`
	Password string `gorm:"not null" json:"password,omitempty" valid:"required~Password is required,minstringlength(6)~Password has to have minimum length of 6 characters"`
	Age      int    `gorm:"not null" json:"age,omitempty" valid:"required~Age is required,range(8|100)~Minimal Age is 8"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hash
	return
}
