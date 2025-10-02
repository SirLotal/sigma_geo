package models

import (
	"fmt"

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
	if err != nil {
		return fmt.Errorf("create (%#v) failed: %w", *prob, err)
	}
	return
}

func (prob *Problem) Update(db *gorm.DB) (err error) {
	err = db.Save(prob).Error
	if err != nil {
		return fmt.Errorf("update to (%#v) failed: %w", *prob, err)
	}
	return
}

func (prob *Problem) Delete(db *gorm.DB) (err error) {
	err = db.Delete(prob).Error
	if err != nil {
		return fmt.Errorf("delete (%#v) failed: %w", *prob, err)
	}
	return
}

func (prob *Problem) Get(db *gorm.DB) (err error) {
	err = db.First(prob).Error
	if err != nil {
		return fmt.Errorf("get (%#v) failed: %w", *prob, err)
	}
	return
}

func (prob *Problem) GetAll(db *gorm.DB, problem_list *[]Problem) (err error) {
	err = db.Where(prob).Find(&problem_list).Error
	if err != nil {
		return fmt.Errorf("get all (%#v) failed: %w", *prob, err)
	}
	return
}
