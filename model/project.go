package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Project struct
type Project struct {
	gorm.Model
	Id          int       `gorm:"not null" json:"id"`
	ProjectName string    `gorm:"not null" json:"projectName"`
	AccountId   int       `gorm:"not null" json:"accountId"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
}
