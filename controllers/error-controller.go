package controllers

import (
	"net/http"

	u "github.com/Raise-the-Voices/raise-the-voices/utils"
)

// NotFound This function responds if the user was wrong with the path
var NotFound = func(w http.ResponseWriter, r *http.Request) {
	resp := u.Message("Error route not found")
	u.Respond(w, resp, 404)
}
