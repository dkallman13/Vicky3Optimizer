package types

import "gorm.io/gorm"

type State struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	name      string `gorm:"column:name"`
	provinces []Province `gorm:"foreignKey:stateId;references:id"`
}

func NewState(stateId int, stateName string) State {
	newstate := State{ID: stateId, name: stateName, provinces: make([]Province, 0)}
	return newstate
}
