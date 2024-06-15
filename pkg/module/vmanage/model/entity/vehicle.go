package entity

import (
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	ID    string `gorm:"type:uuid;default:uuid_generate_v4();primarykey"`
	Title string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (v Vehicle) IsEntity() {}

func (v Vehicle) GetID() string {
	return v.ID
}
