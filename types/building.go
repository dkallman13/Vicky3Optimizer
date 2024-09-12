package types

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	levels int `gorm:"size:64"`
	buildingType string `gorm:"size:32"`
	provinceId int `gorm:"size:64"`
	pms string `gorm:"size:128"`
	inputGoods []Good
	outputGoods []Good
	profit float32 `gorm:"size:20"`
}

func NewBuilding(buildid int, levelnum int, typeof string, provinceid int, pmss string, profits float32) Building{
	newbuilding := Building{id: buildid, levels: levelnum, buildingType: typeof, provinceId: provinceid, pms: pmss, profit: profits, inputGoods: make([]Good, 0), outputGoods: make([]Good, 0)}
	return newbuilding
}