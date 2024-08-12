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