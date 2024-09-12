package initial

import (
	"github.com/dkallman13/Vicky3Optimizer/types"
)
func Migrate(){
	DB.AutoMigrate(types.State{},types.Province{},types.Building{}, types.Good{})
}