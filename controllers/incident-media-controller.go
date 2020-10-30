package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	awsService "github.com/Raise-the-Voices/raise-the-voices/utils/AWS"

	"github.com/Raise-the-Voices/raise-the-voices/middlewares"
	"github.com/Raise-the-Voices/raise-the-voices/models"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/gorilla/mux"
)

// AddIncidentMedia controller request to add a new incident media
var AddIncidentMedia = func(w http.ResponseWriter, r *http.Request) {

	incidentMedia := &models.IncidentMedia{}
	err := json.NewDecoder(r.Body).Decode(incidentMedia) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idincident"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}

	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := incidentMedia.AddIncidentMedia(u.ConvertStringtoUint(vars["idincident"])) //add a incident media
	u.Respond(w, resp, status)
}

// DeleteIncidentMedia controller to delete a media from a incident by idincident-media
var DeleteIncidentMedia = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idincident-media"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}

	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := models.DeleteIncidentMediaByID(u.ConvertStringtoUint(vars["idincident-media"]))
	u.Respond(w, resp, status)
}

// GetIncidentMedia depending on query
var GetIncidentMedia = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()

	switch {
	case (len(urlValues) == 1 && urlValues.Get("idincident") != ""):
		if !u.EvaluationNumberPattern(urlValues.Get("idincident"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetAllMediaByIncidentID(u.ConvertStringtoUint(urlValues.Get("idincident")))
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("idincident-media") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("idincident-media"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetSpecificIncidentMedia(u.ConvertStringtoUint(urlValues.Get("idincident-media")))
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// UploadIncidentFile controller
var UploadIncidentFile = func(w http.ResponseWriter, r *http.Request) {
	var maxUploadSize int64 = 32 << 20

	// FILE Verfication

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
	if ok := u.IsValidFileFormat(filetype); !ok {
		response := u.Message("Error! It is not a permitted format")
		u.Respond(w, response, 500)
		return
	}

	finalName, thumbnailFinalName := u.GenerateFileName(path.Ext(fileheader.Filename))
	fileReader := bytes.NewReader(fileBytes)

	erroraws := awsService.PutObjectToAWSS3(finalName, filetype, fileReader)
	if erroraws != nil {
		log.Println(erroraws)
		response := u.Message("Error! on AWS")
		u.Respond(w, response, 500)
		return
	}

	if u.IsValidImageFileFormat(filetype) {

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

		erroraws := awsService.PutObjectToAWSS3(thumbnailFinalName, filetype, reader)
		if erroraws != nil {
			log.Println(erroraws)
			response := u.Message("Error! on AWS")
			u.Respond(w, response, 500)
			return
		}
	}

	incidentMediaURL := awsService.GenerateProfileS3URL(finalName)

	response := u.Message("the upload was successful")
	response["incident-media-url"] = incidentMediaURL
	u.Respond(w, response, 200)
}
