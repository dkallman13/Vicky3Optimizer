package types

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	Id int `gorm:"primaryKey"`
	Levels int `gorm:"size:64"`
	BuildingType string `gorm:"size:32"`
	ProvinceId int `gorm:"size:64"`
	Pms string `gorm:"size:128"`
	InputGoods []Good
	OutputGoods []Good
	Profit float32 `gorm:"size:20"`
}

func NewBuilding(buildid int, levelnum int, typeof string, provinceid int, pmss string, profits float32) Building{
	newbuilding := Building{Id: buildid, Levels: levelnum, BuildingType: typeof, ProvinceId: provinceid, Pms: pmss, Profit: profits, InputGoods: make([]Good, 0), OutputGoods: make([]Good, 0)}
	return newbuilding
}