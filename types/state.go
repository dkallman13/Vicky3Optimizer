package types

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Id   int `gorm:"primaryKey"`
	Name string `gorm:"size:32"`
	Provinces []Province `gorm:"foreignKey:StateId;references:Id"`
	AirableLand int
	IronCap int
	CoalCap int
	SulfurCap int
	LeadCap int
	WoodCap int
	FishCap int
	WhaleCap int
	OilCap int
	RubberCap int
	HasOpium bool
	HasCotton bool
	HasDye bool
	HasSilk bool
	
}

func NewState(stateId int, stateName string) State {
	newstate := State{Id: stateId, Name: stateName, Provinces: make([]Province, 0)}
	return newstate
}
