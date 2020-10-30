package models

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"
	awsService "github.com/Raise-the-Voices/raise-the-voices/utils/AWS"

	"github.com/jinzhu/gorm"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

// Victim struct
type Victim struct {
	gorm.Model
	Name              string              `json:"name"`
	LegalName         string              `json:"legal_name"`
	Aliases           string              `json:"aliases"`
	PlaceOfBirth      string              `json:"place_of_birth"`
	DateOfBirth       *time.Time          `json:"date_of_birth"`
	CurrentStatus     string              `json:"current_status"`
	Country           string              `json:"country"`
	Gender            string              `json:"gender"`
	LastSeenDate      *time.Time          `json:"last_seen_date"`
	LastSeenPlace     string              `json:"last_seen_place"`
	ProfileImageURL   string              `json:"profile_image_url"`
	Report            Report              `gorm:"foreignkey:VictimID"`
	Incident          []Incident          `gorm:"foreignkey:VictimID"`
	VictimMedia       []VictimMedia       `gorm:"foreignkey:VictimID"`
	VictimTranslation []VictimTranslation `gorm:"foreignkey:VictimID"`
}

// SearchFilter to filter victim for some filters
type SearchFilter struct {
	VictimName  string
	Country     string
	Status      string
	ReportState string
	Sort        string
}

// CreateVictim account db
func (victim *Victim) CreateVictim() (map[string]interface{}, uint) {

	if resp, ok := victim.ValidateVictim(); !ok {
		return resp, 400
	}

	victim.Report.State = "in-review"

	database.GetDB().Create(victim)

	response := u.Message("report has been created")
	response["victim"] = victim
	return response, 201
}

// ValidateVictim Victim Struct
func (victim *Victim) ValidateVictim() (map[string]interface{}, bool) {
	victim.Name = strings.ToLower(victim.Name)
	if victim.Name == "" {
		return u.Message("Name is required"), false
	}

	if victim.Country == "" {
		return u.Message("Country is required"), false
	}

	if len(victim.VictimTranslation) < 1 {
		return u.Message("Victim must have at least 1 translation"), false
	}

	if resp, ok := victim.Report.ValidateReport(); !ok {
		return resp, false
	}

	for i := 0; i < len(victim.VictimTranslation); i++ {
		if resp, ok := victim.VictimTranslation[i].ValidateVictimTranslation(); !ok {
			resp["ERROR Victim Translation number:"] = i + 1
			return resp, false
		}
	}

	// If there is no incident is not need a validation
	// but if there is more one incident each incident must be evaluated
	for i := 0; i < len(victim.Incident); i++ {
		if resp, ok := victim.Incident[i].ValidateIncident(); !ok {
			resp["ERROR incident number:"] = i + 1
			return resp, false
		}
	}

	// If there is no victim media is not need a validation
	// but if there is more one victim media each media must be evaluated
	for i := 0; i < len(victim.VictimMedia); i++ {
		if resp, ok := victim.VictimMedia[i].ValidateVictimMedia(); !ok {
			resp["ERROR victim media number:"] = i + 1
			return resp, false
		}
	}

	return u.Message("Requirement passed"), true
}

// GetSpecificVictim get specific victim
func GetSpecificVictim(idvictim uint, isAuthorizedRequest bool) (map[string]interface{}, uint) {

	specificVictim := &Victim{}

	err := database.GetDB().Table("victims").Where("id = ?", idvictim).First(&specificVictim).Error
	if err != nil {
		response := u.Message("Victim not found")
		return response, 404
	}

	if !isAuthorizedRequest {
		specificVictim.DateOfBirth = nil
	}

	response :=
		map[string]interface{}{
			"victim": specificVictim,
		}
	return response, 200
}

// SearchVictimbyName get victim by name.
func SearchVictimbyName(victimName string, isAuthorizedRequest bool) (map[string]interface{}, uint) {
	victimName = strings.ToLower(victimName)

	specificVictim := &[]Victim{}

	err := database.GetDB().Table("victims").Where("name ILIKE ?", "%"+victimName+"%").Find(&specificVictim).Error
	if err != nil {
		response := u.Message("Victim not found")
		return response, 404
	}
	if !isAuthorizedRequest {
		for i := 0; i < len(*specificVictim); i++ {
			(*specificVictim)[i].DateOfBirth = nil
		}
	}
	response :=
		map[string]interface{}{
			"victim": specificVictim,
		}
	return response, 200
}

// SearchVictimbyFilter multiple filter
func SearchVictimbyFilter(structFilter *SearchFilter, isAuthorizedRequest bool) (map[string]interface{}, uint) {
	victimsList := &[]Victim{}
	var successQuery bool

	queryToDB, argsWhere := validationEmptyFilter(structFilter)
	if victimsList, successQuery = queryWhere((*structFilter).ReportState, (*structFilter).Sort, queryToDB, argsWhere...); !successQuery {
		return u.Message("Some column in the database does not exist"), 404
	}
	if !isAuthorizedRequest {
		for i := 0; i < len(*victimsList); i++ {
			(*victimsList)[i].DateOfBirth = nil
		}
	}
	response :=
		map[string]interface{}{
			"victim": victimsList,
		}
	return response, 200

}

func queryWhere(reportState string, sortBy string, query interface{}, args ...interface{}) (*[]Victim, bool) {
	list := &[]Victim{}
	DBTemp := database.GetDB().Table("victims").Where(query, args...).Order(sortBy)

	// if reportState is different from "published" or "in review"
	// it will return all cases without filtering by report status
	if reportState == "published" || reportState == "in-review" {
		DBTemp = DBTemp.Joins("left join reports on reports.victim_id = victims.id").Where("reports.state = ?", reportState)
	}
	errorVictim := DBTemp.Find(&list).Error
	if errorVictim != nil {
		return nil, false
	}

	return list, true
}

// GetAllVictims get all victims
func GetAllVictims(isAuthorizedRequest bool) (map[string]interface{}, uint) {

	victimsList := &[]Victim{}

	err := database.GetDB().Table("victims").Order("id").Find(&victimsList).Error
	if err != nil {
		response := u.Message("Victims not found")
		return response, 404
	}
	if !isAuthorizedRequest {
		for i := 0; i < len(*victimsList); i++ {
			(*victimsList)[i].DateOfBirth = nil
		}
	}
	response :=
		map[string]interface{}{
			"victims": victimsList,
		}
	return response, 200
}

// GetVictimsByReportState get all victims by report state
func GetVictimsByReportState(reportState string) (map[string]interface{}, uint) {

	victimsList := &[]Victim{}
	database.GetDB().Table("victims").Order("id").Joins("left join reports on reports.victim_id = victims.id").Where("reports.state = ?", reportState).Scan(&victimsList)

	response :=
		map[string]interface{}{
			"victims": victimsList,
		}
	return response, 200
}

// UpdateVictimData from a victim
func (victim *Victim) UpdateVictimData(idVictimtoUpdate uint) (map[string]interface{}, uint) {

	if resp, ok := victim.validateVictimDataToUpdate(); !ok {
		return resp, 400
	}

	victimToUpdate := &Victim{}

	errorUpdateVictim := database.GetDB().Table("victims").Where("id = ?", idVictimtoUpdate).First(&victimToUpdate).Updates(
		Victim{
			Name:          victim.Name,
			LegalName:     victim.LegalName,
			Aliases:       victim.Aliases,
			PlaceOfBirth:  victim.PlaceOfBirth,
			DateOfBirth:   victim.DateOfBirth,
			CurrentStatus: victim.CurrentStatus,
			Country:       victim.Country,
			Gender:        victim.Gender,
			LastSeenDate:  victim.LastSeenDate,
			LastSeenPlace: victim.LastSeenPlace}).Error
	if errorUpdateVictim != nil {
		response := u.Message("Victim not found")
		return response, 404
	}

	response := u.Message("victim have been updated")
	return response, 200
}

// validateVictimDataToUpdate validation victim data
func (victim *Victim) validateVictimDataToUpdate() (map[string]interface{}, bool) {
	victim.Name = strings.ToLower(victim.Name)
	if victim.Name == "" {
		return u.Message("Name is required"), false
	}

	if victim.Country == "" {
		return u.Message("Country is required"), false
	}

	return u.Message("Requirement passed"), true
}

// DownloadPDF returns a pdf file with the number of victims required
func DownloadPDF(victimArrayUint []uint) (map[string]interface{}, uint) {

	arrayVictims := []Victim{}
	for i := 0; i < len(victimArrayUint); i++ {
		report := Report{}
		victim := Victim{}
		victimTranslation := []VictimTranslation{}
		incidentList := []Incident{}
		victimMedia := []VictimMedia{}

		// QUERY Victim
		errorVictim := database.GetDB().Table("victims").Where("id IN (?)", victimArrayUint[i]).First(&victim).Error
		if errorVictim != nil {
			continue
		}

		// QUERY Report
		database.GetDB().Model(&victim).Association("Report").Find(&report)
		victim.Report = report

		// QUERY Victim Translation. Get only English translation.
		database.GetDB().Model(&victim).Where("language = ?", "en").Association("VictimTranslation").Find(&victimTranslation)
		for j := range victimTranslation {
			victim.VictimTranslation = append(victim.VictimTranslation, victimTranslation[j])
		}

		// QUERY Victim Media.
		database.GetDB().Model(&victim).Association("VictimMedia").Find(&victimMedia)
		for j := range victimMedia {
			victim.VictimMedia = append(victim.VictimMedia, victimMedia[j])
		}

		// QUERY Incident
		database.GetDB().Model(&victim).Association("Incident").Find(&incidentList)

		for j := range incidentList {

			incidentTranslation := []IncidentTranslation{}
			incidentMedia := []IncidentMedia{}

			// QUERY Incident Translation. Get only English translation.
			database.GetDB().Model(&incidentList[j]).Where("language = ?", "en").Association("IncidentTranslation").Find(&incidentTranslation)
			for h := range incidentTranslation {
				incidentList[j].IncidentTranslation = append(incidentList[j].IncidentTranslation, incidentTranslation[h])
			}

			// QUERY Incident Media.
			database.GetDB().Model(&incidentList[j]).Association("IncidentMedia").Find(&incidentMedia)
			for h := range incidentMedia {
				incidentList[j].IncidentMedia = append(incidentList[j].IncidentMedia, incidentMedia[h])
			}

			victim.Incident = append(victim.Incident, incidentList[j])
		}

		arrayVictims = append(arrayVictims, victim)
	}

	responsePDF := makerPDF(arrayVictims)
	if responsePDF == "Error when creating the PDF" {
		response := u.Message("Error when creating the PDF")
		return response, 404
	}
	response := u.Message("the pdf file from victim array was created successfully")
	response["PDF"] = makerPDF(arrayVictims)
	return response, 200

}

// DeleteVictimByReport delete a victim because that victim report was deleted
func DeleteVictimByReport(victimID uint) bool {

	victim := &Victim{}
	errorDeleteVictim := database.GetDB().Table("victims").Where("id = ?", victimID).First(&victim).Error
	if errorDeleteVictim != nil {
		return false
	}

	if victim.ProfileImageURL != "" {
		awsService.DeleteObjectAWSS3(victim.ProfileImageURL)
	}

	errorDeleteVictim = database.GetDB().Model(&victim).Delete(&victim).Error
	if errorDeleteVictim != nil {
		return false
	}

	DeleteIncidentByVictimID(victimID)
	DeleteVictimMediaByVictimID(victimID)
	DeleteTranslationsByVictimID(victimID)
	return true
}

// DeleteVictimByID func to delete a victim by her id
func DeleteVictimByID(idVictim uint) (map[string]interface{}, uint) {

	victimToDelete := &Victim{}
	errorDeleteVictim := database.GetDB().Table("victims").Where("id = ?", idVictim).First(&victimToDelete).Error
	if errorDeleteVictim != nil {
		response := u.Message("victim Not Found")
		return response, 404
	}

	if victimToDelete.ProfileImageURL != "" {
		awsService.DeleteObjectAWSS3(victimToDelete.ProfileImageURL)
	}

	errorDeleteVictim = database.GetDB().Model(&victimToDelete).Delete(&victimToDelete).Error
	if errorDeleteVictim != nil {
		response := u.Message("Internal Server Error")
		return response, 500
	}

	DeleteReportByVictimID(victimToDelete.ID)
	DeleteVictimMediaByVictimID(victimToDelete.ID)
	DeleteTranslationsByVictimID(victimToDelete.ID)
	DeleteIncidentByVictimID(victimToDelete.ID)
	response := u.Message("the victim was deleted successfully")
	return response, 200
}

// makerPDF func to create a pdf file.
func makerPDF(arrayText []Victim) string {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	// Row PDF Header
	m.Row(15, func() {
		m.Col(0, func() {
			m.Text("Reports List", props.Text{
				Size:   15,
				Style:  consts.Bold,
				Align:  consts.Center,
				Top:    4,
				Family: "arial",
			})
		})
	})

	// Separator line
	m.Line(1.0)

	for i := 0; i < len(arrayText); i++ {

		// Row the report number and report number line
		m.Row(10, func() {
			m.Col(12, func() {
				m.Text(fmt.Sprintf("Report ID: %d", arrayText[i].Report.ID), props.Text{
					Size:  10,
					Style: consts.Bold,
					Align: consts.Center,
					Top:   4,
				})
			})
		})

		if arrayText[i].ProfileImageURL != "" {
			m.Row(40, func() {
				m.Col(12, func() {
					_ = m.FileImage(u.DownloadFile(arrayText[i].ProfileImageURL), props.Rect{
						Center:  true,
						Percent: 80,
					})
				})
			})
		}

		// Basic Information Table
		BasicHeader := []string{"Basic Information: "}
		m.TableList(BasicHeader, getBasicInformation(&arrayText[i]), props.TableList{
			ContentProp: props.TableListContent{
				GridSizes: []uint{5, 7},
				Family:    "arial",
			},
			HeaderProp: props.TableListContent{
				GridSizes: []uint{3},
				Size:      12.0,
			},
			AlternatedBackground: &color.Color{
				Red:   200,
				Green: 200,
				Blue:  200,
			},
			Align: consts.Left,
			Line:  true,
		})

		// // Separator line
		// m.Line(8.0)

		for j := 0; j < len(arrayText[i].VictimTranslation); j++ {
			if arrayText[i].VictimTranslation[j].Language == "en" {

				// General Information Table
				generalHeader := []string{"General Information: "}
				m.TableList(generalHeader, getGeneralInformation(&arrayText[i].VictimTranslation[j]), props.TableList{
					ContentProp: props.TableListContent{
						GridSizes: []uint{5, 7},
						Family:    "arial",
					},
					HeaderProp: props.TableListContent{
						GridSizes: []uint{3},
						Size:      12.0,
					},
					AlternatedBackground: &color.Color{
						Red:   200,
						Green: 200,
						Blue:  200,
					},
					HeaderContentSpace: 1.0,
					Align:              consts.Left,
					Line:               true,
				})

			}
		}

		// // Separator line
		// m.Line(8.0)

		// Victim Media Table
		// victimMediaHeader := []string{"Victim Media: "}
		// m.TableList(victimMediaHeader, getVictimMedia(&arrayText[i].VictimMedia), props.TableList{
		// 	ContentProp: props.TableListContent{
		// 		GridSizes: []uint{5, 7},
		// 		Family:    "arial",
		// 	},
		// 	HeaderProp: props.TableListContent{
		// 		GridSizes: []uint{3},
		// 		Size:      12.0,
		// 	},
		// 	AlternatedBackground: &color.Color{
		// 		Red:   200,
		// 		Green: 200,
		// 		Blue:  200,
		// 	},
		// 	Align: consts.Left,
		// 	Line:  true,
		// })

		// // Separator line
		// m.Line(8.0)

		// Incident Array
		for j := 0; j < len(arrayText[i].Incident); j++ {

			// Row the Incident number
			m.Row(10, func() {
				m.Col(12, func() {
					m.Text(fmt.Sprintf("Incident #: %d", j+1), props.Text{
						Size:  10,
						Style: consts.Bold,
						Align: consts.Center,
						Top:   4,
					})
				})
			})

			// Incident Information Table
			incidentHeader := []string{"Incident Information: "}
			m.TableList(incidentHeader, getIncident(&arrayText[i].Incident[j]), props.TableList{
				ContentProp: props.TableListContent{
					GridSizes: []uint{5, 7},
					Family:    "arial",
				},
				HeaderProp: props.TableListContent{
					GridSizes: []uint{3},
					Size:      12.0,
				},
				AlternatedBackground: &color.Color{
					Red:   200,
					Green: 200,
					Blue:  200,
				},
				Align: consts.Left,
				Line:  true,
			})

			// // Separator line
			// m.Line(8.0)

			for k := 0; k < len(arrayText[i].Incident[j].IncidentTranslation); k++ {
				if arrayText[i].Incident[j].IncidentTranslation[k].Language == "en" {

					// Incident Description Table
					incidentDescriptionHeader := []string{"Description: "}
					m.TableList(incidentDescriptionHeader, getIncidentDescription(&arrayText[i].Incident[j].IncidentTranslation[k]), props.TableList{
						ContentProp: props.TableListContent{
							GridSizes: []uint{5, 7},
							Family:    "arial",
						},
						HeaderProp: props.TableListContent{
							GridSizes: []uint{3},
							Size:      12.0,
						},
						AlternatedBackground: &color.Color{
							Red:   200,
							Green: 200,
							Blue:  200,
						},
						Align: consts.Left,
						Line:  true,
					})
				}
			}

			// // Separator line
			// m.Line(8.0)

			// // Incident Media Table
			// incidentMediaHeader := []string{"Incident Media: "}
			// m.TableList(incidentMediaHeader, getIncidentMedia(&arrayText[i].Incident[j].IncidentMedia), props.TableList{
			// 	ContentProp: props.TableListContent{
			// 		GridSizes: []uint{5, 7},
			// 		Family:    "arial",
			// 	},
			// 	HeaderProp: props.TableListContent{
			// 		GridSizes: []uint{3},
			// 		Size:      12.0,
			// 	},
			// 	AlternatedBackground: &color.Color{
			// 		Red:   200,
			// 		Green: 200,
			// 		Blue:  200,
			// 	},
			// 	Align: consts.Left,
			// 	Line:  true,
			// })

		}

		// Separator line
		m.Line(15.0)

	}

	PDFBuffer, err := m.Output()
	if err != nil {
		// fmt.Println("Could not save PDF:", err)
		return "Error when creating the PDF"
	}

	stringBase64PDF := base64.StdEncoding.EncodeToString(PDFBuffer.Bytes())

	return stringBase64PDF
}

func getBasicInformation(victim *Victim) [][]string {

	contents := [][]string{}
	contents = append(contents, []string{"Name: ", victim.Name})
	contents = append(contents, []string{"Legal Name: ", victim.LegalName})
	contents = append(contents, []string{"Aliases: ", victim.Aliases})
	contents = append(contents, []string{"Place of birth: ", victim.PlaceOfBirth})
	contents = append(contents, []string{"Date of birth: ", getDateFieldData(victim.DateOfBirth)})
	contents = append(contents, []string{"Current status: ", victim.CurrentStatus})
	contents = append(contents, []string{"Country: ", victim.Country})
	contents = append(contents, []string{"Gender: ", victim.Gender})
	contents = append(contents, []string{"Last seen place: ", victim.LastSeenPlace})
	contents = append(contents, []string{"Last seen date: ", getDateFieldData(victim.LastSeenDate)})

	return contents
}

func getGeneralInformation(translation *VictimTranslation) [][]string {

	contents := [][]string{}
	contents = append(contents, []string{"Nationality: ", translation.Nationality})
	contents = append(contents, []string{"Health status: ", translation.HealthStatus})
	contents = append(contents, []string{"Health issues: ", translation.HealthIssues})
	contents = append(contents, []string{"Languages spoken: ", translation.LanguagesSpoken})
	contents = append(contents, []string{"Profession: ", translation.Profession})
	contents = append(contents, []string{"About the victim: ", translation.AboutTheVictim})
	contents = append(contents, []string{"Additional information: ", translation.AdditionalInformation})

	return contents
}

func getVictimMedia(media *[]VictimMedia) [][]string {
	contents := [][]string{}
	for i := 0; i < len(*media); i++ {
		contents = append(contents, []string{"URL: ", (*media)[i].MediaURL})
	}

	return contents
}

func getIncident(incident *Incident) [][]string {

	contents := [][]string{}
	contents = append(contents, []string{"Date of incident: ", incident.DateOfIncident.Format("02-01-2006")})
	contents = append(contents, []string{"Location: ", incident.Location})
	contents = append(contents, []string{"Is disappearance: ", strconv.FormatBool(incident.IsDisappearance)})

	return contents
}

func getIncidentDescription(description *IncidentTranslation) [][]string {

	contents := [][]string{}
	contents = append(contents, []string{"Narrative of Incident: ", description.NarrativeOfIncident})
	contents = append(contents, []string{"Current status summary: ", description.CurrentStatusSummary})

	return contents
}

func getIncidentMedia(media *[]IncidentMedia) [][]string {
	contents := [][]string{}
	for i := 0; i < len(*media); i++ {
		contents = append(contents, []string{"URL: ", (*media)[i].MediaURL})
	}

	return contents
}

func validationEmptyFilter(structFilter *SearchFilter) (string, []interface{}) {
	var queryString string
	var argsWhere []interface{}

	if (*structFilter).VictimName != "" {
		queryString = `name ILIKE ?`

		argsWhere = append(argsWhere, "%"+(*structFilter).VictimName+"%")
	}

	if (*structFilter).Status != "all" {
		if queryString != "" {
			queryString += ` AND current_status = ?`
		} else {
			queryString = `current_status = ?`
		}

		argsWhere = append(argsWhere, (*structFilter).Status)
	}

	if (*structFilter).Country != "all" {
		if queryString != "" {
			queryString += ` AND country = ?`
		} else {
			queryString = `country = ?`
		}

		argsWhere = append(argsWhere, (*structFilter).Country)
	}

	return queryString, argsWhere
}

func getDateFieldData(data *time.Time) string {
	if data == nil {
		return "unknown"
	}
	return data.Format("02-01-2006")
}
