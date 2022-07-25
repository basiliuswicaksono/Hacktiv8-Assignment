package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt *time.Time `json:"creted_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
