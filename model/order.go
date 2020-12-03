package model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User User `gorm: "ForeignKey: id"`
	ProjectId int
	Status int8 `gorm:"default:0"`
}
