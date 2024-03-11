package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Email       string    `gorm:"not null,unique" json:"email"`
	IsSuspended bool      `gorm:"default:false" json:"isSuspended"`
	Teachers    []Teacher `gorm:"many2many:student_teachers;" json:"teachers"`
}