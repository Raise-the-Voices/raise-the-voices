package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Raise-the-Voices/raise-the-voices/middlewares"
	"github.com/Raise-the-Voices/raise-the-voices/models"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"
	"github.com/gorilla/mux"
)

// AddOptionController controller request to add a new option
var AddOptionController = func(w http.ResponseWriter, r *http.Request) {

	option := &models.Options{}
	err := json.NewDecoder(r.Body).Decode(option) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "admin", withRespond); !ok {
		return
	}

	resp, status := option.AddOption() //add an option
	u.Respond(w, resp, status)
}

// DeleteOptionController controller to delete a option
var DeleteOptionController = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["optionid"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "admin", withRespond); !ok {
		return
	}

	resp, status := models.DeleteOption(u.ConvertStringtoUint(vars["optionid"]))
	u.Respond(w, resp, status)
}

// GetAllOptionsController depending on query
var GetAllOptionsController = func(w http.ResponseWriter, r *http.Request) {
	resp, status := models.GetAllOptions()
	u.Respond(w, resp, status)
}
