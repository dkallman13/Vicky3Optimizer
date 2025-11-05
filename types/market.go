package types

import "gorm.io/gorm"

type Market struct{
	gorm.Model
	Id int `gorm:"primaryKey"`
	Name string
	Goods []Good `gorm:"many2many:MarketGoods;"`
	Countries []Country `gorm:"many2many:MarketCountries;"`
}