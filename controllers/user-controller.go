package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/Raise-the-Voices/raise-the-voices/middlewares"
	"github.com/Raise-the-Voices/raise-the-voices/models"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/gorilla/mux"
)

// CreateUser from a administrator user
var CreateUser = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "admin", withRespond); !ok {
		return
	}

	resp, status := user.UserCreation() //Create account
	u.Respond(w, resp, status)
}

// UpdateUserPassword controller
var UpdateUserPassword = func(w http.ResponseWriter, r *http.Request) {

	password := &models.Password{}
	err := json.NewDecoder(r.Body).Decode(password) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := password.UpdatePassword() //update data from a user
	u.Respond(w, resp, status)
}

// UpdateUserData controller
var UpdateUserData = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	vars := mux.Vars(r)

	if !u.EvaluationNumberPattern(vars["iduser"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := user.UpdateUser(u.ConvertStringtoUint(vars["iduser"])) //update data from a user
	u.Respond(w, resp, status)
}

// GetUsers depending on query
var GetUsers = func(w http.ResponseWriter, r *http.Request) {

	parsedURL, _ := url.Parse(r.URL.String())
	urlValues := parsedURL.Query()
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "admin", withRespond); !ok {
		return
	}

	switch {
	case (len(urlValues) == 1 && urlValues.Get("iduser") != ""):

		if !u.EvaluationNumberPattern(urlValues.Get("iduser"), false) {
			response := u.Message("Params Error")
			u.Respond(w, response, 400)
			return
		}

		resp, status := models.GetSpecificUser(u.ConvertStringtoUint(urlValues.Get("iduser")))
		u.Respond(w, resp, status)
		break
	case (len(urlValues) == 1 && urlValues.Get("name") != ""):

		resp, status := models.SearchUserbyName(urlValues.Get("name"))
		u.Respond(w, resp, status)
		break
	case len(urlValues) == 0:
		resp, status := models.GetAllUsers()
		u.Respond(w, resp, status)
		break
	default:
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
	}

}

// DeleteUser controller to delete a user by its id
var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "admin", withRespond); !ok {
		return
	}

	if !u.EvaluationNumberPattern(vars["iduser"], false) {
		response := u.Message("Params Error")
		u.Respond(w, response, 400)
		return
	}

	resp, status := models.DeleteUserByID(u.ConvertStringtoUint(vars["iduser"]))
	u.Respond(w, resp, status)
}

// Authenticate a user
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	resp, status := models.Login(strings.ToLower(user.Email), user.Password)
	u.Respond(w, resp, status)
}

// RefreshToken user token refresh controller
var RefreshToken = func(w http.ResponseWriter, r *http.Request) {
	withRespond := true
	if ok := middlewares.AuthenticationAPI(w, r, "all", withRespond); !ok {
		return
	}

	resp, status := models.RefreshUserToken(middlewares.GetClaims())
	u.Respond(w, resp, status)
}

// PasswordEmail get email from a user to validate it and send him a email to reset password
var PasswordEmail = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	resp, status := user.GenerateEmailToResetPassword()
	u.Respond(w, resp, status)
}

// ResetToken Get token and validate it
var ResetToken = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if vars["token"] == "" {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	resp, status := models.ValidateResetToken(vars["token"])
	u.Respond(w, resp, status)
}

// ResetPassword from a user.
var ResetPassword = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message("Invalid request"), 400)
		return
	}

	resp, status := user.ResetPasswordFromUser()
	u.Respond(w, resp, status)
}
