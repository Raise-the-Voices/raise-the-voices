package controllers

import (
	"bytes"
	"encoding/json"
	"path"

	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	awsService "github.com/Raise-the-Voices/raise-the-voices/utils/AWS"

	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/Raise-the-Voices/raise-the-voices/database"
	"github.com/Raise-the-Voices/raise-the-voices/middlewares"
	"github.com/Raise-the-Voices/raise-the-voices/models"

	"github.com/gorilla/mux"
)

// CreateVictimReport controller
var CreateVictimReport = func(w http.ResponseWriter, r *http.Request) {

	victim := &models.Victim{}
	err := json.NewDecoder(r.Body).Decode(victim) //decode the request body into struct and failed if any error occur

	if err != nil {
		log.Println("Error: " + err.Error())
		u.Respond(w, u.Message("Invalid request"+err.Error()), 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := victim.CreateVictim() //Create victim report
	u.Respond(w, resp, status)
}

// GetVictims depending on query
var GetVictims = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()
	withRespond := false
	isAuthorizedRequest := middlewares.AuthenticationAPI(w, r, "all", withRespond)

	switch {
	case (len(urlValues) == 1 && urlValues.Get("idvictim") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("idvictim"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		if models.IsReportInReview(u.ConvertStringtoUint(urlValues.Get("idvictim"))) {
			if !isAuthorizedRequest {
				response := u.Message("You are not authorized to access.")
				u.Respond(w, response, 400)
				return
			}
		}

		resp, status := models.GetSpecificVictim(u.ConvertStringtoUint(urlValues.Get("idvictim")), isAuthorizedRequest)

		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("victim-name") != ""):

		resp, status := models.SearchVictimbyName(urlValues.Get("victim-name"), isAuthorizedRequest)
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("report-state") != ""):

		resp, status := models.GetVictimsByReportState(urlValues.Get("report-state"))
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 5 && urlValues.Get("country") != "" &&
		urlValues.Get("status") != "" && urlValues.Get("report-state") != "" &&
		urlValues.Get("sort") != ""):

		structFilter := &models.SearchFilter{
			VictimName:  urlValues.Get("victim-name"),
			Country:     urlValues.Get("country"),
			Status:      urlValues.Get("status"),
			ReportState: urlValues.Get("report-state"),
			Sort:        urlValues.Get("sort")}

		resp, status := models.SearchVictimbyFilter(structFilter, isAuthorizedRequest)
		u.Respond(w, resp, status)
		break
	case len(urlValues) == 0:
		resp, status := models.GetAllVictims(isAuthorizedRequest)
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// UpdateVictimData controller
var UpdateVictimData = func(w http.ResponseWriter, r *http.Request) {

	victim := &models.Victim{}
	err := json.NewDecoder(r.Body).Decode(victim) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idvictim"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := victim.UpdateVictimData(u.ConvertStringtoUint(vars["idvictim"])) //update data from a victim
	u.Respond(w, resp, status)
}

// DeleteVictim controller to delete a victim by her id
var DeleteVictim = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idvictim"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := models.DeleteVictimByID(u.ConvertStringtoUint(vars["idvictim"]))
	u.Respond(w, resp, status)
}

// GetVictimsPDF get pdf file on victims
var GetVictimsPDF = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	victimsArray := parsedURL.Query()["idvictim"]

	switch {
	case (len(victimsArray) > 0):
		var VictimsUintArray []uint
		// []string
		for i := 0; i < len(victimsArray); i++ {
			if !u.EvaluationNumberPattern(victimsArray[i], false) {
				response := u.Message("Params Error")
				u.Respond(w, response, 400)
				return
			}
			VictimsUintArray = append(VictimsUintArray, u.ConvertStringtoUint(victimsArray[i]))
		}

		resp, status := models.DownloadPDF(VictimsUintArray)
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// UploadProfilePhoto controller
var UploadProfilePhoto = func(w http.ResponseWriter, r *http.Request) {
	var maxUploadSize int64 = 32 << 20
	vars := mux.Vars(r)
	specificVictim := &models.Victim{}

	if !u.EvaluationNumberPattern(vars["idvictim"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}

	// Verfication FILE

	// 32MB is the default used by FormFile
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		response := u.Message("Error! body form-part exceeds the allowed limit")
		u.Respond(w, response, 500)
		return
	}

	file, fileheader, err := r.FormFile("myfile")
	if err != nil {
		response := u.Message("Error! formatting the file")
		u.Respond(w, response, 500)
		return
	}
	defer file.Close()

	fileSize := fileheader.Size
	if u.IsGreaterObjectSizeThanLimit(fileSize, maxUploadSize) {
		response := u.Message("Error! The Object exceeds the allowed limit")
		u.Respond(w, response, 500)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		response := u.Message("Error! Reading the entire object")
		u.Respond(w, response, 500)
		return
	}

	filetype := http.DetectContentType(fileBytes)
	if ok := u.IsValidImageFileFormat(filetype); !ok {
		response := u.Message("Error! It is not a permitted format")
		u.Respond(w, response, 500)
		return
	}

	finalName, thumbnailFinalName := u.GenerateFileName(path.Ext(fileheader.Filename))

	// Resize
	src, thumbFormatType, errDecode := u.ResizeImage(file)
	if errDecode != nil {
		response := u.Message("Error! resizing the image")
		u.Respond(w, response, 500)
		return
	}

	buff, errEncode := u.EncodeImage(thumbFormatType, src)
	if errEncode != nil {
		response := u.Message("Error! Encoding the image")
		u.Respond(w, response, 500)
		return
	}

	// convert bufferImage to reader
	reader := bytes.NewReader(buff.Bytes())
	fileReader := bytes.NewReader(fileBytes)

	// DB

	err = database.GetDB().Table("victims").Where("id = ?", u.ConvertStringtoUint(vars["idvictim"])).First(&specificVictim).Error
	if err != nil {
		response := u.Message("Victim not found")
		u.Respond(w, response, 404)
		return
	}

	// AWS
	if specificVictim.ProfileImageURL != "" {
		awsService.DeleteObjectAWSS3(specificVictim.ProfileImageURL)
	}

	erroraws := awsService.PutObjectToAWSS3(finalName, filetype, fileReader)
	if erroraws != nil {
		log.Println(erroraws)
		response := u.Message("Error! on AWS")
		u.Respond(w, response, 500)
		return
	}

	erroraws = awsService.PutObjectToAWSS3(thumbnailFinalName, filetype, reader)
	if erroraws != nil {
		log.Println(erroraws)
		response := u.Message("Error! on AWS")
		u.Respond(w, response, 500)
		return
	}

	ProfileURL := awsService.GenerateProfileS3URL(finalName)

	// DB Assign image path in database.
	specificVictim.ProfileImageURL = ProfileURL
	database.GetDB().Save(&specificVictim)

	response := u.Message("the upload was successful")
	u.Respond(w, response, 200)
}

// DeleteProfilePhoto controller
var DeleteProfilePhoto = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	specificVictim := &models.Victim{}

	if !u.EvaluationNumberPattern(vars["idvictim"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}

	err := database.GetDB().Table("victims").Where("id = ?", u.ConvertStringtoUint(vars["idvictim"])).First(&specificVictim).Error
	if err != nil {
		response := u.Message("Victim not found")
		u.Respond(w, response, 404)
		return
	}

	if specificVictim.ProfileImageURL != "" {
		resultDelete, errDelete := awsService.DeleteObjectAWSS3(specificVictim.ProfileImageURL)
		if errDelete != nil {
			response := u.Message(resultDelete)
			u.Respond(w, response, 404)
			return
		}

		// Assign image path in database.
		specificVictim.ProfileImageURL = ""
		database.GetDB().Save(&specificVictim)

		response := u.Message(resultDelete)
		u.Respond(w, response, 200)
		return
	}

	response := u.Message("the url field is empty")
	u.Respond(w, response, 404)
}
