package types

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	ID   int
	Name string `gorm:"size:32"`
}

func NewState(stateId int, stateName string) State {
	newstate := State{ID: stateId, Name: stateName}
	return newstate
}
