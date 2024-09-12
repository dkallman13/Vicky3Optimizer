package types

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	name string `gorm:"size:64"`
	states []State 
}
func NewCountry(ID int, countryname string) Country{
	newcountry := Country{id :ID, name: countryname,  states: make([]State, 0)}
	return newcountry
}