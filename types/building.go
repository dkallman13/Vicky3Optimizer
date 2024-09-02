package types

import (
)

type Building struct {
	id int
	levels int
	buildingType string
	stateId int
	pms string
	inputGoods []Good
	outputGoods []Good
	profit float32
}

func NewBuilding(buildid int, levelnum int, typeof string, stateid int, pmss string, profits float32) Building{
	newbuilding := Building{id: buildid, levels: levelnum, buildingType: typeof, stateId: stateid, pms: pmss, profit: profits, inputGoods: make([]Good, 0), outputGoods: make([]Good, 0)}
	return newbuilding
}