package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Raise-the-Voices/raise-the-voices/middlewares"
	"github.com/Raise-the-Voices/raise-the-voices/models"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/gorilla/mux"
)

// AddIncidentTranslationController controller request
var AddIncidentTranslationController = func(w http.ResponseWriter, r *http.Request) {

	incidentTranslation := &models.IncidentTranslation{}
	err := json.NewDecoder(r.Body).Decode(incidentTranslation) //decode the request body into struct and failed if any error occur

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

	resp, status := incidentTranslation.AddIncidentTranslation(u.ConvertStringtoUint(vars["idincident"])) //add a incident translation
	u.Respond(w, resp, status)
}

// UpdateIncidentTranslation controller
var UpdateIncidentTranslation = func(w http.ResponseWriter, r *http.Request) {

	incidentTranslation := &models.IncidentTranslation{}
	err := json.NewDecoder(r.Body).Decode(incidentTranslation) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idincident-translation"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}

	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := incidentTranslation.UpdateIncidentTranslation(u.ConvertStringtoUint(vars["idincident-translation"])) //update data from a incident translation
	u.Respond(w, resp, status)
}

// GetIncidentTranslations depending on query
var GetIncidentTranslations = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()

	switch {
	case (len(urlValues) == 1 && urlValues.Get("idincident") != ""):
		if !u.EvaluationNumberPattern(urlValues.Get("idincident"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetAllTranslationsByIncidentID(u.ConvertStringtoUint(urlValues.Get("idincident")))
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("idincident-translation") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("idincident-translation"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetSpecificIncidentTranslation(u.ConvertStringtoUint(urlValues.Get("idincident-translation")))
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// DeleteIncidentTranslation controller to delete a translation from a incident by id incident-translation
var DeleteIncidentTranslation = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idincident-translation"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := models.DeleteIncidentTranslationByID(u.ConvertStringtoUint(vars["idincident-translation"]))
	u.Respond(w, resp, status)
}
