package types

import "gorm.io/gorm"

type Pop struct {
	gorm.Model
	id int
	jobType string
	workforceSize int
	dependantsSize int
	workplaceId int
}
func NewPop(popid int, job string, workforce int, dependants int, workid int) Pop{
	newpop := Pop{id: popid, jobType: job, workforceSize: workforce, dependantsSize: dependants, workplaceId: workid}
	return newpop
}