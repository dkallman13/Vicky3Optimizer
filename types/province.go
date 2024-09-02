package types

import (
)

type Province struct{
	id int
	stateId int
	buildings []Building
}
func NewProvince(ID int, stateid int) Province{
	newprovince := Province{id :ID, stateId: stateid,  buildings: make([]Building, 0)}
	return newprovince
}