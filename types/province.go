package types

import "gorm.io/gorm"

type Province struct{
	gorm.Model
	Id int `gorm:"primaryKey"`
	StateId int `gorm:"size:64"`
	Buildings []Building `gorm:"foreignKey:provinceId;references:id"`
}
func NewProvince(ID int, stateid int) Province{
	newprovince := Province{Id :ID, StateId: stateid,  Buildings: make([]Building, 0)}
	return newprovince
}