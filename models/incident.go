package models

import (
	"time"

	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/jinzhu/gorm"
)

// Incident struct
type Incident struct {
	gorm.Model
	VictimID            uint                  `json:"victimid"`
	DateOfIncident      time.Time             `json:"date_of_incident"`
	Location            string                `json:"location"`
	Type                string                `json:"type"`
	IsDisappearance     bool                  `json:"is_disappearance"`
	IncidentTranslation []IncidentTranslation `gorm:"foreignkey:IncidentID"`
	IncidentMedia       []IncidentMedia       `gorm:"foreignkey:IncidentID"`
}

// ValidateIncident Incident Struct
func (incident *Incident) ValidateIncident() (map[string]interface{}, bool) {

	if incident.DateOfIncident.IsZero() {
		return u.Message("Date of Incident is required"), false
	}

	if incident.Location == "" {
		return u.Message("Incident Location is required"), false
	}

	if incident.IsDisappearance != true && incident.IsDisappearance != false {
		return u.Message("IsDisappearance is required"), false
	}

	// If there is no incident Narrative is not need a validation
	// but if there is more one incident Narrative each narrative must be evaluated
	for i := 0; i < len(incident.IncidentTranslation); i++ {
		if resp, ok := incident.IncidentTranslation[i].ValidateIncidentTranslation(); !ok {
			resp["ERROR incident Narrative number:"] = i + 1
			return resp, false
		}
	}

	// If there is no incident media is not need a validation
	// but if there is more one incident media each media must be evaluated
	for i := 0; i < len(incident.IncidentMedia); i++ {
		if resp, ok := incident.IncidentMedia[i].ValidateIncidentMedia(); !ok {
			resp["ERROR incident media number:"] = i + 1
			return resp, false
		}
	}

	return u.Message("Requirement passed"), true
}

// AddIncident method to add a incident related with victimID
func (incident *Incident) AddIncident(victimID uint) (map[string]interface{}, uint) {
	if resp, ok := incident.ValidateIncident(); !ok {
		return resp, 400
	}

	if resp, ok := ValidationExistVictimID(victimID); !ok {
		return resp, 404
	}

	incident.VictimID = victimID
	database.GetDB().Create(incident)

	response := u.Message("Incident has been added")
	response["incident"] = incident
	return response, 201
}

// UpdateVictimIncident update incident data from a victim
func (incident *Incident) UpdateVictimIncident(idincident uint) (map[string]interface{}, uint) {

	if resp, ok := incident.ValidateIncident(); !ok {
		return resp, 400
	}

	incidentToUpdate := &Incident{}

	errorUpdateIncident := database.GetDB().Table("incidents").Where("id = ?", idincident).First(&incidentToUpdate).Updates(
		Incident{
			DateOfIncident: incident.DateOfIncident,
			Location:       incident.Location,
			Type:           incident.Type}).Error

	if errorUpdateIncident != nil {
		response := u.Message("Incident not found or was deleted")
		return response, 404
	}

	incidentToUpdate.IsDisappearance = incident.IsDisappearance

	database.GetDB().Save(&incidentToUpdate)

	response := u.Message("Incident have been updated")
	return response, 200
}

// GetAllIncidentsByVictimID get all incidents from a victim
func GetAllIncidentsByVictimID(idVictim uint) (map[string]interface{}, uint) {

	victimIncidentList := &[]Incident{}

	err := database.GetDB().Table("incidents").Where("victim_id = ?", idVictim).Order("id").Find(&victimIncidentList).Error
	if err != nil {
		response := u.Message("Incidents not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"incidents": victimIncidentList,
		}
	return response, 200
}

// DeleteIncidentByVictimID delete a incident by victim id
func DeleteIncidentByVictimID(victimID uint) bool {

	incident := &[]Incident{}
	errorDeleteIncident := database.GetDB().Table("incidents").Where("victim_id = ?", victimID).Find(&incident).Delete(&incident).Error
	if errorDeleteIncident != nil {
		return false
	}

	for i := 0; i < len(*incident); i++ {
		DeleteIncidentMediaByIncident((*incident)[i].ID)
		DeleteIncidentTranslationsByIncidentID((*incident)[i].ID)
	}

	return true
}

// Falta incluir borrar media y translation

// DeleteIncidentByID delete a translation by its id
func DeleteIncidentByID(idincident uint) (map[string]interface{}, uint) {

	victimIncident := &Incident{}
	errorDeleteIncident := database.GetDB().Table("incidents").Where("id = ?", idincident).First(&victimIncident).Delete(&victimIncident).Error
	if errorDeleteIncident != nil {
		return u.Message("Incident not Found"), 404
	}

	DeleteIncidentMediaByIncident(idincident)
	DeleteIncidentTranslationsByIncidentID(idincident)

	return u.Message("Incident was deleted successfully"), 200
}

// ValidationExistIncidentID Function to validate if incidentID exist in incident must to move to a share folder
func ValidationExistIncidentID(incidentID uint) (map[string]interface{}, bool) {
	existIncident := &Incident{}
	errorExistIncident := database.GetDB().Table("incidents").Where("id = ?", incidentID).Find(&existIncident).Error
	if errorExistIncident != nil {
		return u.Message("Incident ID not Exist"), false
	}

	return u.Message("Requirement passed"), true
}
