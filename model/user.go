package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `grom:"type:varchar(20);not null"`
	Telephone string `grom:"varchar(110);not null;unique"`
	Password string `gorm:"size:255;not null"`
}