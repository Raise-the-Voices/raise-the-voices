package models

import (
	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/jinzhu/gorm"
)

// VictimTranslation struct
type VictimTranslation struct {
	gorm.Model
	VictimID              uint   `json:"victimid"`
	Language              string `json:"language"`
	Nationality           string `json:"nationality"`
	HealthStatus          string `json:"health_status"`
	HealthIssues          string `json:"health_issues"`
	LanguagesSpoken       string `json:"languagues_spoken"`
	Profession            string `json:"profession"`
	AboutTheVictim        string `json:"about_the_victim"`
	AdditionalInformation string `json:"additional_information"`
}

// ValidateVictimTranslation ValidateVictimTranslation Struct
func (victimTranslation *VictimTranslation) ValidateVictimTranslation() (map[string]interface{}, bool) {

	if victimTranslation.Language == "" || len(victimTranslation.Language) != 2 {
		return u.Message("Language is required"), false
	}

	return u.Message("Requirement passed"), true
}

// AddVictimTranslation method to add a victim translation related with victimID
func (victimTranslation *VictimTranslation) AddVictimTranslation(victimID uint) (map[string]interface{}, uint) {
	if resp, ok := victimTranslation.ValidateVictimTranslation(); !ok {
		return resp, 400
	}

	if resp, ok := ValidationExistVictimID(victimID); !ok {
		return resp, 404
	}

	victimTranslation.VictimID = victimID
	database.GetDB().Create(victimTranslation)

	response := u.Message("victim Translation has been added")
	response["victim_translation"] = victimTranslation
	return response, 201
}

// DeleteTranslationsByVictimID delete victim translations by victim id
func DeleteTranslationsByVictimID(victimID uint) bool {

	victimTranslations := &[]VictimTranslation{}
	errorDeleteTranslations := database.GetDB().Table("victim_translations").Where("victim_id = ?", victimID).Find(&victimTranslations).Delete(&victimTranslations).Error
	if errorDeleteTranslations != nil {
		return false
	}

	return true
}

// DeleteTranslationByID delete a translation by its id
func DeleteTranslationByID(idtranslation uint) (map[string]interface{}, uint) {

	victimTranslation := &VictimTranslation{}
	errorDeleteTranslation := database.GetDB().Table("victim_translations").Where("id = ?", idtranslation).First(&victimTranslation).Delete(&victimTranslation).Error
	if errorDeleteTranslation != nil {
		return u.Message("Translation not Found"), 404
	}

	return u.Message("translation was deleted successfully"), 200
}

// ValidationExistVictimID Function to validate if VictimID exist in victim must to move to a share folder
func ValidationExistVictimID(victimID uint) (map[string]interface{}, bool) {
	existVictim := &Victim{}
	errorExistVictim := database.GetDB().Table("victims").Where("id = ?", victimID).Find(&existVictim).Error
	if errorExistVictim != nil {
		return u.Message("Victim ID not Exist"), false
	}

	return u.Message("Requirement passed"), true
}

// UpdateVictimTranslation update translation from a victim
func (victimTranslation *VictimTranslation) UpdateVictimTranslation(idTranslationtoUpdate uint) (map[string]interface{}, uint) {

	if resp, ok := victimTranslation.ValidateVictimTranslation(); !ok {
		return resp, 400
	}

	translationToUpdate := &VictimTranslation{}

	errorUpdateVictim := database.GetDB().Table("victim_translations").Where("id = ?", idTranslationtoUpdate).First(&translationToUpdate).Updates(
		VictimTranslation{
			Language:              victimTranslation.Language,
			Nationality:           victimTranslation.Nationality,
			HealthStatus:          victimTranslation.HealthStatus,
			HealthIssues:          victimTranslation.HealthIssues,
			LanguagesSpoken:       victimTranslation.LanguagesSpoken,
			Profession:            victimTranslation.Profession,
			AboutTheVictim:        victimTranslation.AboutTheVictim,
			AdditionalInformation: victimTranslation.AdditionalInformation}).Error

	if errorUpdateVictim != nil {
		response := u.Message("Victim Translation not found or was deleted")
		return response, 404
	}

	response := u.Message("Victim Translation have been updated")
	return response, 200
}

// GetAllTranslationsByVictimID get all translations from a victim
func GetAllTranslationsByVictimID(idVictim uint) (map[string]interface{}, uint) {

	victimTranslationList := &[]VictimTranslation{}

	err := database.GetDB().Table("victim_translations").Where("victim_id = ?", idVictim).Order("id").Find(&victimTranslationList).Error
	if err != nil {
		response := u.Message("Translations not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"translations": victimTranslationList,
		}
	return response, 200
}

// GetSpecificTranslation get specific translation from a victim
func GetSpecificTranslation(idtranslation uint) (map[string]interface{}, uint) {

	specificTranslation := &VictimTranslation{}

	err := database.GetDB().Table("victim_translations").Where("id = ?", idtranslation).First(&specificTranslation).Error
	if err != nil {
		response := u.Message("Translation not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"translation": specificTranslation,
		}
	return response, 200
}
