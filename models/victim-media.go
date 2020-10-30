package models

import (
	"time"

	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"
	awsService "github.com/Raise-the-Voices/raise-the-voices/utils/AWS"

	"github.com/jinzhu/gorm"
)

// VictimMedia struct
type VictimMedia struct {
	gorm.Model
	VictimID    uint      `json:"victimid"`
	DateOfMedia time.Time `json:"date_of_media"`
	MediaURL    string    `json:"mediaurl"`
	Tag         string    `json:"tag"`
}

// ValidateVictimMedia Victim Media Struct
func (victimMedia *VictimMedia) ValidateVictimMedia() (map[string]interface{}, bool) {

	if victimMedia.DateOfMedia.IsZero() {
		return u.Message("Date of Victim Media is required"), false
	}

	if victimMedia.MediaURL == "" {
		return u.Message("URL Victim Media is required"), false
	}

	return u.Message("Requirement passed"), true
}

// AddVictimMedia method to add a victim media related with victimID
func (victimMedia *VictimMedia) AddVictimMedia(victimID uint) (map[string]interface{}, uint) {
	if resp, ok := victimMedia.ValidateVictimMedia(); !ok {
		return resp, 400
	}

	if resp, ok := ValidationExistVictimID(victimID); !ok {
		return resp, 404
	}

	victimMedia.VictimID = victimID
	database.GetDB().Create(victimMedia)

	response := u.Message("victim media has been added")
	response["victim_media"] = victimMedia
	return response, 201
}

// DeleteVictimMediaByID delete a victim media by its id
func DeleteVictimMediaByID(idvictimmedia uint) (map[string]interface{}, uint) {

	victimMedia := &VictimMedia{}
	errorDeleteVictimMedia := database.GetDB().Table("victim_media").Where("id = ?", idvictimmedia).First(&victimMedia).Delete(&victimMedia).Error
	if errorDeleteVictimMedia != nil {
		return u.Message("Media not Found"), 404
	}

	if victimMedia.MediaURL != "" {
		awsService.DeleteObjectAWSS3(victimMedia.MediaURL)
	}

	return u.Message("Media was deleted successfully"), 200
}

// DeleteVictimMediaByVictimID delete victim media by victim id
func DeleteVictimMediaByVictimID(victimID uint) bool {

	victimMedias := &[]VictimMedia{}
	errorDeleteMedias := database.GetDB().Table("victim_media").Where("victim_id = ?", victimID).Find(&victimMedias).Delete(&victimMedias).Error
	if errorDeleteMedias != nil {
		return false
	}

	for i := 0; i < len(*victimMedias); i++ {
		if (*victimMedias)[i].MediaURL != "" {
			awsService.DeleteObjectAWSS3((*victimMedias)[i].MediaURL)
		}
	}

	return true
}

// GetAllMediaByVictimID get all media from a victim
func GetAllMediaByVictimID(idVictim uint, authorizedRequest bool) (map[string]interface{}, uint) {

	victimMediaList := &[]VictimMedia{}

	victimDB := database.GetDB().Table("victim_media").Where("victim_id = ?", idVictim)

	if !authorizedRequest {
		victimDB = victimDB.Where("tag <> ?", "documents")
	}

	err := victimDB.Order("id").Find(&victimMediaList).Error
	if err != nil {
		response := u.Message("Medias not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"medias": victimMediaList,
		}
	return response, 200
}

// GetSpecificMedia get specific media from a victim
func GetSpecificMedia(idvictimmedia uint, authorizedRequest bool) (map[string]interface{}, uint) {

	specificMedia := &VictimMedia{}

	victimDB := database.GetDB().Table("victim_media").Where("id = ?", idvictimmedia)

	if !authorizedRequest {
		victimDB = victimDB.Where("tag <> ?", "documents")
	}

	err := victimDB.First(&specificMedia).Error
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
