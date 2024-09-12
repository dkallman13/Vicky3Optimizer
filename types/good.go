package types

import "gorm.io/gorm"

type Good struct {
	gorm.Model
	id int
	name string
	amount float32
	price float32
}
func NewGood(goodid int, goodname string, goodamount float32, goodprice float32) Good{
	newgood := Good{id :goodid, name: goodname, amount: goodamount, price: goodprice}
	return newgood
}