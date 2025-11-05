package model

import (
	"github.com/dkallman13/Vicky3Optimizer/types"
)
func Migrate(){
	DB.AutoMigrate(types.State{},types.Province{},types.Building{}, types.Good{}, types.Country{}, types.Pop{},types.Market{}, types.Qualifications{}, types.SaveFile{})
}