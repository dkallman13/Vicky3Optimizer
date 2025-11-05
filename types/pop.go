package types

import "gorm.io/gorm"

type Pop struct {
	gorm.Model
	Id int `gorm:"primaryKey"`
	JobType string
	WorkforceSize int
	DependantsSize int
	WorkplaceId int 
	//qualifications:
	//0 academics
	//1 aristocrats
	//2 bureaucrats
	//3 capitalists
	//4 clergymen
	//5 clerks
	//6 engineers
	//7 farmers
	//8 laborers (always 100%)
	//9 machinists
	//10 officers
	//11 soldiers (always 100%)
	//12 shopkeepers
	qualifications []Qualifications 
	jobSatisfaction float64
	soL int
	wealth int
}

func NewPop(popid int, job string, workforce int, dependants int, workid int, qual []Qualifications, jobsat float64, sol int, Wealth int) Pop{
	newpop := Pop{Id: popid, JobType: job, WorkforceSize: workforce, DependantsSize: dependants, WorkplaceId: workid, qualifications: qual,jobSatisfaction: jobsat, soL: sol, wealth: Wealth}
	return newpop
}