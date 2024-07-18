package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	CompanyName string
	TradeName   string
	Site        string
}
