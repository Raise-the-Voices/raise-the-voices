package models

import (
	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/jinzhu/gorm"
)

// IncidentTranslation struct
type IncidentTranslation struct {
	gorm.Model
	IncidentID           uint   `json:"incidentid"`
	Language             string `json:"language"`
	NarrativeOfIncident  string `json:"narrative_of_incident"`
	CurrentStatusSummary string `json:"current_status_summary"`
}

// ValidateIncidentTranslation Incident Struct
func (incidentTranslation *IncidentTranslation) ValidateIncidentTranslation() (map[string]interface{}, bool) {

	if incidentTranslation.Language == "" || len(incidentTranslation.Language) != 2 {
		return u.Message("Language is required"), false
	}

	if incidentTranslation.NarrativeOfIncident == "" {
		return u.Message("NarrativeOfIncident is required"), false
	}

	return u.Message("Requirement passed"), true
}

// AddIncidentTranslation method to add a incident translation related with idincident
func (incidentTranslation *IncidentTranslation) AddIncidentTranslation(idincident uint) (map[string]interface{}, uint) {
	if resp, ok := incidentTranslation.ValidateIncidentTranslation(); !ok {
		return resp, 400
	}

	if resp, ok := ValidationExistIncidentID(idincident); !ok {
		return resp, 404
	}

	incidentTranslation.IncidentID = idincident
	database.GetDB().Create(incidentTranslation)

	response := u.Message("incident Translation has been added")
	response["incident_translation"] = incidentTranslation
	return response, 201
}

// UpdateIncidentTranslation update translation from a incident
func (incidentTranslation *IncidentTranslation) UpdateIncidentTranslation(idincidentTranslation uint) (map[string]interface{}, uint) {

	if resp, ok := incidentTranslation.ValidateIncidentTranslation(); !ok {
		return resp, 400
	}

	incidentTranslationToUpdate := &IncidentTranslation{}

	errorUpdateIncidentTranslation := database.GetDB().Table("incident_translations").Where("id = ?", idincidentTranslation).First(&incidentTranslationToUpdate).Updates(
		IncidentTranslation{
			Language:             incidentTranslation.Language,
			NarrativeOfIncident:  incidentTranslation.NarrativeOfIncident,
			CurrentStatusSummary: incidentTranslation.CurrentStatusSummary}).Error

	if errorUpdateIncidentTranslation != nil {
		response := u.Message("Incident Translation not found or was deleted")
		return response, 404
	}

	response := u.Message("Incident Translation have been updated")
	return response, 200
}

// GetAllTranslationsByIncidentID get all translations from a incident
func GetAllTranslationsByIncidentID(idincident uint) (map[string]interface{}, uint) {

	incidentTranslationList := &[]IncidentTranslation{}

	err := database.GetDB().Table("incident_translations").Where("incident_id = ?", idincident).Order("id").Find(&incidentTranslationList).Error
	if err != nil {
		response := u.Message("Incidents not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"translations": incidentTranslationList,
		}
	return response, 200
}

// GetSpecificIncidentTranslation get specific translation from a incident
func GetSpecificIncidentTranslation(idtranslation uint) (map[string]interface{}, uint) {

	specificTranslation := &IncidentTranslation{}

	err := database.GetDB().Table("incident_translations").Where("id = ?", idtranslation).First(&specificTranslation).Error
	if err != nil {
		response := u.Message("Incident Translation not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"translation": specificTranslation,
		}
	return response, 200
}

// DeleteIncidentTranslationsByIncidentID delete incident translations by incident ID
func DeleteIncidentTranslationsByIncidentID(incidentID uint) bool {

	incidentTranslations := &[]IncidentTranslation{}
	errorDeleteTranslations := database.GetDB().Table("incident_translations").Where("incident_id = ?", incidentID).Find(&incidentTranslations).Delete(&incidentTranslations).Error
	if errorDeleteTranslations != nil {
		return false
	}

	return true
}

// DeleteIncidentTranslationByID delete a translation by its idincident-translation
func DeleteIncidentTranslationByID(idIncidentTranslation uint) (map[string]interface{}, uint) {

	incidentTranslation := &IncidentTranslation{}
	errorDeleteTranslation := database.GetDB().Table("incident_translations").Where("id = ?", idIncidentTranslation).First(&incidentTranslation).Delete(&incidentTranslation).Error
	if errorDeleteTranslation != nil {
		return u.Message("Translation not Found"), 404
	}

	return u.Message("translation was deleted successfully"), 200
}
