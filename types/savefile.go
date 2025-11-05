package types

import (
	"gorm.io/gorm"

)

type SaveFile struct {
	gorm.Model
	SaveId int `gorm:"primaryKey"`
	PlayerCountryId int
	PlayerCountryName string
	Countries []Country `gorm:"foreignKey:Id;references:Id"`
	Pops []Pop `gorm:"foreignKey:Id;references:Id"`
}