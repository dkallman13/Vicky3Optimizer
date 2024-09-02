package types

import (
)

type Country struct {
	id int
	name string
	states []State
}
func NewCountry(ID int, countryname string) Country{
	newcountry := Country{id :ID, name: countryname,  states: make([]State, 0)}
	return newcountry
}