package types

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Id   int `gorm:"primaryKey"`
	Name string `gorm:"size:32"`
	Provinces []Province `gorm:"foreignKey:StateId;references:Id"`
}

func NewState(stateId int, stateName string) State {
	newstate := State{Id: stateId, Name: stateName, Provinces: make([]Province, 0)}
	return newstate
}
