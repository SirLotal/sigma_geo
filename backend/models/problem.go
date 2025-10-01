package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Text       string  `json:"text"`
	Sol        *string `json:"sol"`
	Section    string  `json:"section"`
	Theme      *string `json:"theme"`
	Proposer   *string `json:"proposer"`
	Difficulty *uint   `json:"difficulty"`
	Source     *string `json:"source"`
	Comment    *string `json:"comment"`
	Picture    *string `json:"picture"`

	SubProjectName *string `json:"subproject_name"`
	VariantName    *string `json:"variant_name"`
	Number         *string `json:"number"`
}

func (prob *Problem) Create(db *gorm.DB) (err error) {
	err = db.Save(prob).Error
	return
}

func (prob *Problem) Update(db *gorm.DB) (err error) {
	err = db.Save(prob).Error
	return
}

func (prob *Problem) Delete(db *gorm.DB) (err error) {
	err = db.Delete(prob).Error
	return
}

func (prob *Problem) Get(db *gorm.DB) (err error) {
	err = db.First(prob).Error
	return
}
