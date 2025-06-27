package models

import (
	"time"
)

type User struct {
	Id          uint    `json:"id" gorm:"primarykey;autoIncrement"` 
	Name        string    `json:"name" gorm:"not null"`
	CreatedDate time.Time `json:"createdDate" gorm:"not null; default:CURRENT_TIMESTAMP"`
	UpdatedDate time.Time `json:"updatedDate" gorm:"not null; default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
