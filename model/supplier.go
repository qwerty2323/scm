package model

import (
	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Name string
	Inn int
	Kpp int
	PaymentInfo string
}
