package types

import "gorm.io/gorm"

type Province struct{
	gorm.Model
	id int `gorm:"primaryKey"`
	stateId int `gorm:"size:64"`
	buildings []Building `gorm:"foreignKey:provinceId;references:id"`
}
func NewProvince(ID int, stateid int) Province{
	newprovince := Province{id :ID, stateId: stateid,  buildings: make([]Building, 0)}
	return newprovince
}