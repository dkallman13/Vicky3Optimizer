package types

import "gorm.io/gorm"

type Country struct {
	gorm.Model
	Id int `gorm:"primaryKey"`
	Name string `gorm:"size:64"`
	Provinces []Province `gorm:"foreignKey:CountryId;references:Id"`
}
func NewCountry(ID int, countryname string) Country{
	newcountry := Country{Id :ID, Name: countryname,  Provinces: make([]Province, 0)}
	return newcountry
}