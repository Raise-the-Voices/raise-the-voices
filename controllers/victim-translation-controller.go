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

// AddVictimTranslationController controller request
var AddVictimTranslationController = func(w http.ResponseWriter, r *http.Request) {

	victimTranslation := &models.VictimTranslation{}
	err := json.NewDecoder(r.Body).Decode(victimTranslation) //decode the request body into struct and failed if any error occur

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

	resp, status := victimTranslation.AddVictimTranslation(u.ConvertStringtoUint(vars["idvictim"])) //add a victim translation
	u.Respond(w, resp, status)
}

// UpdateTranslationData controller
var UpdateTranslationData = func(w http.ResponseWriter, r *http.Request) {

	victimTranslation := &models.VictimTranslation{}
	err := json.NewDecoder(r.Body).Decode(victimTranslation) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idvictim-translation"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := victimTranslation.UpdateVictimTranslation(u.ConvertStringtoUint(vars["idvictim-translation"])) //update data from a victim translation
	u.Respond(w, resp, status)
}

// GetVictimTranslations depending on query
var GetVictimTranslations = func(w http.ResponseWriter, r *http.Request) {

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

		resp, status := models.GetAllTranslationsByVictimID(u.ConvertStringtoUint(urlValues.Get("idvictim")))
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("idvictim-translation") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("idvictim-translation"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetSpecificTranslation(u.ConvertStringtoUint(urlValues.Get("idvictim-translation")))
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// DeleteTranslation controller to delete a translation from a victim by id translation
var DeleteTranslation = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["idvictim-translation"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := models.DeleteTranslationByID(u.ConvertStringtoUint(vars["idvictim-translation"]))
	u.Respond(w, resp, status)
}
