package types

import "gorm.io/gorm"

type Good struct {
	gorm.Model
	Id int `gorm:"primaryKey"`
	Name string
	Amount float64
	Price float64
}
func NewGood(goodid int, goodname string, goodamount float64, goodprice float64) Good{
	newgood := Good{Id :goodid, Name: goodname, Amount: goodamount, Price: goodprice}
	return newgood
}