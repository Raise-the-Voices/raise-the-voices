package models

import (
	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"
	"github.com/jinzhu/gorm"
)

// Options struct
type Options struct {
	gorm.Model
	Title string `json:"title"`
	Group string `json:"group"`
}

// AddOption method to add a new option
func (option *Options) AddOption() (map[string]interface{}, uint) {

	database.GetDB().Create(option)

	response := u.Message("option has been added")
	response["option"] = option
	return response, 201
}

// DeleteOption delete a option
func DeleteOption(optionID uint) (map[string]interface{}, uint) {

	option := &Options{}
	err := database.GetDB().Table("options").Where("id = ?", optionID).Delete(&option).Error
	if err != nil {
		return u.Message("option not found"), 404
	}

	return u.Message("option was deleted successfully"), 200
}

// GetAllOptions get all options
func GetAllOptions() (map[string]interface{}, uint) {

	optionList := &[]Options{}
	err := database.GetDB().Table("options").Order("id").Find(&optionList).Error
	if err != nil {
		response := u.Message("option not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"options-list": optionList,
		}
	return response, 200
}
