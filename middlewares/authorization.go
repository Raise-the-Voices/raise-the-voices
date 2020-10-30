package middlewares

import (
	"net/http"

	"github.com/Raise-the-Voices/raise-the-voices/models"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"
)

// AutorizationAPI permission level user
func AutorizationAPI(w http.ResponseWriter, permissionType string, claims *models.Claims, withRespond bool) bool {

	if permissionType == claims.UserRole || permissionType == "all" {
		return true
	}

	if withRespond {
		u.Respond(w, u.Message("You are not authorized to access."), 403)
	}

	return false
}
