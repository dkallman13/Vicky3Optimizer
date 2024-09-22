package types

import "gorm.io/gorm"

type Pop struct {
	gorm.Model
	Id int
	JobType string
	WorkforceSize int
	DependantsSize int
	WorkplaceId int
}
func NewPop(popid int, job string, workforce int, dependants int, workid int) Pop{
	newpop := Pop{Id: popid, JobType: job, WorkforceSize: workforce, DependantsSize: dependants, WorkplaceId: workid}
	return newpop
}