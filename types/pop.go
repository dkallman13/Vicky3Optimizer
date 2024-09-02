package types

import (
)

type Pop struct {
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