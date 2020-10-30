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

// AddIncidentController controller request
var AddIncidentController = func(w http.ResponseWriter, r *http.Request) {

	newIncident := &models.Incident{}
	err := json.NewDecoder(r.Body).Decode(newIncident) //decode the request body into struct and failed if any error occur

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

	resp, status := newIncident.AddIncident(u.ConvertStringtoUint(vars["idvictim"])) //add a incident
	u.Respond(w, resp, status)
}

// UpdateIncidentData controller
var UpdateIncidentData = func(w http.ResponseWriter, r *http.Request) {

	updateIncident := &models.Incident{}
	err := json.NewDecoder(r.Body).Decode(updateIncident) //decode the request body into struct and failed if any error occur

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

	resp, status := updateIncident.UpdateVictimIncident(u.ConvertStringtoUint(vars["idincident"])) //update data from a incident
	u.Respond(w, resp, status)
}

// GetIncidents depending on query
var GetIncidents = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()

	switch {
	case (len(urlValues) == 1 && urlValues.Get("idvictim") != ""):
		if !u.EvaluationNumberPattern(urlValues.Get("idvictim"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		if models.IsReportInReview(u.ConvertStringtoUint(urlValues.Get("idvictim"))) {
			withRespond := true
			if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
				return
			}
		}

		resp, status := models.GetAllIncidentsByVictimID(u.ConvertStringtoUint(urlValues.Get("idvictim")))
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// DeleteIncident controller to delete a incident from a victim by id incident
var DeleteIncident = func(w http.ResponseWriter, r *http.Request) {
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

	resp, status := models.DeleteIncidentByID(u.ConvertStringtoUint(vars["idincident"]))
	u.Respond(w, resp, status)
}
