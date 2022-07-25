package models

type Photo struct {
	GormModel
	Title     string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Caption   string `json:"caption"`
	Photo_url string `gorm:"not null" json:"photo_url" valid:"required~Photo_url is required"`
	UserID    uint

	User *User // check lagi
}
