package types

import "gorm.io/gorm"

type Qualifications struct {
	gorm.Model
	id int `gorm:"primaryKey"`
	popId int `gorm:"foreignKey:popId"`
	jobId int64
	qualifiedPops float64
}
func NewQualifications(popid int, jobid int64, qualifiedpops float64) Qualifications{
	newqualifications := Qualifications{popId: popid, jobId: jobid, qualifiedPops: qualifiedpops}
	return newqualifications
}