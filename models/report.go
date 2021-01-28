package models

import (
	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/jinzhu/gorm"
)

// Report structs
type Report struct {
	gorm.Model
	VictimID          uint   `json:"victimid"`
	State             string `json:"state"`
	NameOfReporter    string `json:"name_of_reporter"`
	EmailOfReporter   string `json:"email_of_reporter"`
	Discovery         string `json:"discovery"`
	IsDirectTestimony bool   `json:"is_direct_testimony"`
}

// ValidateReport Struct
func (report *Report) ValidateReport() (map[string]interface{}, bool) {

	if report.NameOfReporter == "" {
		return u.Message("name of reporter is required"), false
	}

	if report.EmailOfReporter == "" {
		return u.Message("email of reporter is required"), false
	}

	if report.Discovery == "" {
		return u.Message("discovery is required"), false
	}
	// if the isdirecttestimony field does not come by request the default report object will make false.
	if report.IsDirectTestimony != true && report.IsDirectTestimony != false {
		return u.Message("IsDirectTestimony is required"), false
	}

	return u.Message("Requirement passed"), true
}

// UpdateReport update data of a specify report.
func (report *Report) UpdateReport(idReport uint) (map[string]interface{}, uint) {

	if resp, ok := report.ValidateReport(); !ok {
		return resp, 400
	}

	reportToUpdate := &Report{}

	errorUpdateReport := database.GetDB().Table("reports").Where("id = ?", idReport).First(&reportToUpdate).Updates(
		Report{
			NameOfReporter:    report.NameOfReporter,
			EmailOfReporter:   report.EmailOfReporter,
			Discovery:         report.Discovery,
			IsDirectTestimony: report.IsDirectTestimony}).Error

	if errorUpdateReport != nil {
		response := u.Message("report not found or was deleted")
		return response, 404
	}
	reportToUpdate.IsDirectTestimony = report.IsDirectTestimony
	database.GetDB().Save(&reportToUpdate)

	response := u.Message("Report have been updated")
	return response, 200
}

// ModifyState change report state
func ModifyState(idReport uint) (map[string]interface{}, uint) {

	reportState := &Report{Model: gorm.Model{ID: idReport}}

	errorGetVictimReport := database.GetDB().First(reportState).Error
	if errorGetVictimReport != nil {
		response := u.Message("report Not Found")
		return response, 404
	}

	switch reportState.State {
	case "in-review":
		database.GetDB().Model(&reportState).Update("state", "published")
		break
	case "published":
		database.GetDB().Model(&reportState).Update("state", "in-review")
		break
	}

	response := u.Message("report state modified")
	return response, 200
}

// GetReportList get report list depending on query
func GetReportList(state string, offset uint, limit uint) (map[string]interface{}, uint) {

	reportsList := &[]Report{}
	var err error

	switch state {
	case "all":
		err = database.GetDB().Table("reports").Order("id").Offset(offset).Limit(limit).Find(&reportsList).Error
		break
	case "in-review":
		err = database.GetDB().Table("reports").Order("id").Where("state = ?", state).Offset(offset).Limit(limit).Find(&reportsList).Error
		break
	case "published":
		err = database.GetDB().Table("reports").Order("id").Where("state = ?", state).Offset(offset).Limit(limit).Find(&reportsList).Error
		break
	default:
		response := u.Message("Params:State Error")
		return response, 400
	}

	if err != nil {
		response := u.Message("Error report list")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"reports": reportsList,
		}
	return response, 200
}

// GetSpecificReport get specific report
func GetSpecificReport(idreport uint) (map[string]interface{}, uint) {

	specificReport := &Report{}

	err := database.GetDB().Table("reports").Where("id = ?", idreport).First(&specificReport).Error
	if err != nil {
		response := u.Message("Report not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"report": specificReport,
		}
	return response, 200
}

// GetSpecificReportByVictimID get specific report by victim-id
func GetSpecificReportByVictimID(victimID uint) (map[string]interface{}, uint) {

	specificReport := &Report{}

	err := database.GetDB().Table("reports").Where("victim_id = ?", victimID).First(&specificReport).Error
	if err != nil {
		response := u.Message("Report not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"report": specificReport,
		}
	return response, 200
}

// GetAllReport get all report
func GetAllReport() (map[string]interface{}, uint) {

	reportsList := &[]Report{}

	err := database.GetDB().Table("reports").Order("id").Find(&reportsList).Error
	if err != nil {
		response := u.Message("Reports not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"reports": reportsList,
		}
	return response, 200
}

// DeleteReportByID func to delete a report by id
func DeleteReportByID(idReport uint) (map[string]interface{}, uint) {

	reportState := &Report{}
	errorDeleteReport := database.GetDB().Table("reports").Where("id = ?", idReport).First(&reportState).Delete(&reportState).Error
	if errorDeleteReport != nil {
		response := u.Message("report Not Found")
		return response, 404
	}

	DeleteVictimByReport(reportState.VictimID)
	response := u.Message("the report was deleted successfully")
	return response, 200
}

// DeleteReportByVictimID func to delete a report by victim id
func DeleteReportByVictimID(idVictim uint) bool {

	report := &Report{}
	errorDeleteReport := database.GetDB().Table("reports").Where("victim_id = ?", idVictim).First(&report).Delete(&report).Error
	if errorDeleteReport != nil {
		return false
	}

	return true
}

// IsReportInReview evaluate if report is in review
func IsReportInReview(idvictim uint) bool {
	specificReport := &Report{}
	err := database.GetDB().Table("reports").Where("victim_id = ?", idvictim).First(&specificReport).Error
	if err != nil {
		return false
	}
	if specificReport.State == "in-review" {
		return true
	}
	return false
}
