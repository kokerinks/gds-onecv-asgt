package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Email    string    `gorm:"not null,unique" json:"email"`
	Students []Student `gorm:"many2many:student_teachers;" json:"students"`
}