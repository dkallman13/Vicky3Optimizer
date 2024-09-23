package types

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	Id int `gorm:"primaryKey"`
	Levels int `gorm:"size:64"`
	BuildingType string `gorm:"size:32"`
	ProvinceId int `gorm:"size:64"`
	Pms string `gorm:"size:128"`
	InputGoods []Good `gorm:"many2many:BuildingIGoods;"`
	OutputGoods []Good `gorm:"many2many:BuildingOGoods;"`
	Profit float64 `gorm:"size:20"`
	WorkerPops []Pop `gorm:"foreignKey:WorkplaceId;references:Id"`
}

func NewBuilding(buildid int, levelnum int, typeof string, provinceid int, pmss string, profits float64) Building{
	newbuilding := Building{Id: buildid, Levels: levelnum, BuildingType: typeof, ProvinceId: provinceid, Pms: pmss, Profit: profits, InputGoods: make([]Good, 0), OutputGoods: make([]Good, 0), WorkerPops: make([]Pop, 0)}
	return newbuilding
}