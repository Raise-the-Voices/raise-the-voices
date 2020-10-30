package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/Raise-the-Voices/raise-the-voices/middlewares"
	"github.com/Raise-the-Voices/raise-the-voices/models"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"
	awsService "github.com/Raise-the-Voices/raise-the-voices/utils/AWS"

	"github.com/gorilla/mux"
)

// AddVictimMediaController controller request to add a new victim media
var AddVictimMediaController = func(w http.ResponseWriter, r *http.Request) {

	victimMedia := &models.VictimMedia{}
	err := json.NewDecoder(r.Body).Decode(victimMedia) //decode the request body into struct and failed if any error occur

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

	resp, status := victimMedia.AddVictimMedia(u.ConvertStringtoUint(vars["idvictim"])) //add a victim media
	u.Respond(w, resp, status)
}

// DeleteVictimMediaController controller to delete a media from a victim by id media
var DeleteVictimMediaController = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idvictimmedia"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := models.DeleteVictimMediaByID(u.ConvertStringtoUint(vars["idvictimmedia"]))
	u.Respond(w, resp, status)
}

// GetVictimMediaController depending on query
var GetVictimMediaController = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()
	withRespond := false
	authorizedRequest := middlewares.AuthenticationAPI(w, r, "all", withRespond)

	switch {
	case (len(urlValues) == 1 && urlValues.Get("idvictim") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("idvictim"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		if models.IsReportInReview(u.ConvertStringtoUint(urlValues.Get("idvictim"))) {
			if !authorizedRequest {
				u.Respond(w, u.Message("You are not authorized to access."), 403)
				return
			}
		}

		resp, status := models.GetAllMediaByVictimID(u.ConvertStringtoUint(urlValues.Get("idvictim")), authorizedRequest)
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("idvictimmedia") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("idvictimmedia"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetSpecificMedia(u.ConvertStringtoUint(urlValues.Get("idvictimmedia")), authorizedRequest)
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// UploadVictimFile controller
var UploadVictimFile = func(w http.ResponseWriter, r *http.Request) {
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

	victimMediaURL := awsService.GenerateProfileS3URL(finalName)

	response := u.Message("the upload was successful")
	response["victim-media-url"] = victimMediaURL
	u.Respond(w, response, 200)
}
