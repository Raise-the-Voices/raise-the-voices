package models

import (
	"time"

	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"
	awsService "github.com/Raise-the-Voices/raise-the-voices/utils/AWS"

	"github.com/jinzhu/gorm"
)

// IncidentMedia struct
type IncidentMedia struct {
	gorm.Model
	IncidentID  uint      `json:"incidentid"`
	DateOfMedia time.Time `json:"date_of_media"`
	MediaURL    string    `json:"mediaurl"`
	Tag         string    `json:"tag"`
}

// ValidateIncidentMedia Incident Media Struct
func (incidentMedia *IncidentMedia) ValidateIncidentMedia() (map[string]interface{}, bool) {

	if incidentMedia.DateOfMedia.IsZero() {
		return u.Message("Date of Incident Media is required"), false
	}

	if incidentMedia.MediaURL == "" {
		return u.Message("URL Incident Media is required"), false
	}

	return u.Message("Requirement passed"), true
}

// AddIncidentMedia method to add a incident media related with incidentID
func (incidentMedia *IncidentMedia) AddIncidentMedia(incidentID uint) (map[string]interface{}, uint) {
	if resp, ok := incidentMedia.ValidateIncidentMedia(); !ok {
		return resp, 400
	}

	if resp, ok := ValidationExistIncidentID(incidentID); !ok {
		return resp, 404
	}

	incidentMedia.IncidentID = incidentID
	database.GetDB().Create(incidentMedia)

	response := u.Message("incident media has been added")
	response["incident_media"] = incidentMedia
	return response, 201
}

// GetAllMediaByIncidentID get all media from a incident
func GetAllMediaByIncidentID(idincident uint) (map[string]interface{}, uint) {

	incidentMediaList := &[]IncidentMedia{}

	err := database.GetDB().Table("incident_media").Where("incident_id = ?", idincident).Order("id").Find(&incidentMediaList).Error
	if err != nil {
		response := u.Message("Medias not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"medias": incidentMediaList,
		}
	return response, 200
}

// GetSpecificIncidentMedia get specific media from a incident
func GetSpecificIncidentMedia(idincidentmedia uint) (map[string]interface{}, uint) {

	specificMedia := &IncidentMedia{}

	err := database.GetDB().Table("incident_media").Where("id = ?", idincidentmedia).First(&specificMedia).Error
	if err != nil {
		response := u.Message("media not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"media": specificMedia,
		}
	return response, 200
}

// DeleteIncidentMediaByID delete a incident media by its id
func DeleteIncidentMediaByID(idincidentmedia uint) (map[string]interface{}, uint) {

	incidentMedia := &IncidentMedia{}
	errorDeleteIncidentMedia := database.GetDB().Table("incident_media").Where("id = ?", idincidentmedia).First(&incidentMedia).Delete(&incidentMedia).Error
	if errorDeleteIncidentMedia != nil {
		return u.Message("Media not Found"), 404
	}

	if incidentMedia.MediaURL != "" {
		awsService.DeleteObjectAWSS3(incidentMedia.MediaURL)
	}

	return u.Message("Media was deleted successfully"), 200
}

// DeleteIncidentMediaByIncident delete incident Medias by incident id
func DeleteIncidentMediaByIncident(incidentID uint) bool {

	incidentMedias := &[]IncidentMedia{}
	errorDeleteMedias := database.GetDB().Table("incident_media").Where("incident_id = ?", incidentID).Find(&incidentMedias).Delete(&incidentMedias).Error
	if errorDeleteMedias != nil {
		return false
	}

	for i := 0; i < len(*incidentMedias); i++ {
		if (*incidentMedias)[i].MediaURL != "" {
			awsService.DeleteObjectAWSS3((*incidentMedias)[i].MediaURL)
		}
	}

	return true
}
