package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Raise-the-Voices/raise-the-voices/middlewares"
	"github.com/Raise-the-Voices/raise-the-voices/models"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/gorilla/mux"
)

// ModifyReportState modify report state from "in-review" to "published" or viceversa
var ModifyReportState = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idreport"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := models.ModifyState(u.ConvertStringtoUint(vars["idreport"]))
	u.Respond(w, resp, status)
}

// UpdateReportData controller
var UpdateReportData = func(w http.ResponseWriter, r *http.Request) {

	updateReport := &models.Report{}
	err := json.NewDecoder(r.Body).Decode(updateReport) //decode the request body into struct and failed if any error occur

	if err != nil {
		log.Println("Error: " + err.Error())
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idreport"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := updateReport.UpdateReport(u.ConvertStringtoUint(vars["idreport"])) //update data from a incident
	u.Respond(w, resp, status)
}

// GetReport depending on query
var GetReport = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	switch {
	case (urlValues.Get("state") != "" && urlValues.Get("offset") != "" && urlValues.Get("limit") != "" && len(urlValues) == 3):

		if !u.EvaluationNumberPattern(urlValues.Get("offset"), true) ||
			!u.EvaluationNumberPattern(urlValues.Get("limit"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetReportList(strings.ToLower(urlValues.Get("state")), u.ConvertStringtoUint(urlValues.Get("offset")), u.ConvertStringtoUint(urlValues.Get("limit")))
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("idreport") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("idreport"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetSpecificReport(u.ConvertStringtoUint(urlValues.Get("idreport")))
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("victimID") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("victimID"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetSpecificReportByVictimID(u.ConvertStringtoUint(urlValues.Get("victimID")))
		u.Respond(w, resp, status)
		break
	case len(urlValues) == 0:
		resp, status := models.GetAllReport()
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// DeleteReport controller to delete a victim report by id
var DeleteReport = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idreport"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := models.DeleteReportByID(u.ConvertStringtoUint(vars["idreport"]))
	u.Respond(w, resp, status)
}
