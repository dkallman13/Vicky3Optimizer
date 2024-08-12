package types

import (
)

type SaveFile struct {
	PlayerCountryId int
	PlayerCountryName string
	Countries []Country
	Goods []Good
	Pops []Pop
}