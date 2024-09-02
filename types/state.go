package types

import (
)


type State struct{
	id int
	name string
	provinces []Province 
}
func NewState(stateId int, stateName string) State{
	newstate := State{id :stateId, name: stateName,  provinces: make([]Province, 0)}
	return newstate
}