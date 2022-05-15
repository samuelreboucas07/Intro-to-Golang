package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model        //Add id, createAt, updateAt....
	Name       string `json:"name" validate:"nonzero"`
	CPF        string `json:"cpf" validate:"len=11,regexp=^[0-9]*$"`
	RG         string `json:"rg" validate:"len=9,regexp=^[0-9]*$"`
}

func ValidateStudent(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}
	return nil
}