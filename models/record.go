package models

import (
	"gorm.io/gorm"
)

// Define the Record model
type Record struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(100);"`
	Email    string `gorm:"column:email;type:varchar(100);"`
	Phone    string `gorm:"column:phone;type:varchar(20);"`
	Location string `gorm:"column:location;type:varchar(100);"`
}
