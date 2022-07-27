package models

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" valid:"required~Message is required"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`

	User  *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photo *Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
