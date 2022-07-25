package models

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" valid:"required~Message is required"`
	UserID  uint
	PhotoID uint

	User  *User // check lagi
	Photo *User // check lagi
}
