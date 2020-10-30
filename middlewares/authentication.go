package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/Raise-the-Voices/raise-the-voices/models"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/dgrijalva/jwt-go"
)

var tk = &models.Claims{}

// AuthenticationAPI middleware function to evaluate if user token is valid
func AuthenticationAPI(w http.ResponseWriter, r *http.Request, permissionType string, withRespond bool) bool {

	// JWT
	tokenHeader := r.Header.Get("Authorization") //Grab the token from the header
	if tokenHeader == "" {
		if withRespond {
			u.Respond(w, u.Message("Missing auth token"), 403) //Token is missing, returns with error code 403 Unauthorized
		}
		return false
	}

	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		if withRespond {
			u.Respond(w, u.Message("Invalid/Malformed auth token"), 403)
		}
		return false
	}

	tokenPart := splitted[1] //Grab the token part, what we are truly interested in

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})

	if err != nil { //Malformed token, returns with http code 403 as usual
		if withRespond {
			u.Respond(w, u.Message("Malformed or Expired authentication token"), 403)
		}
		return false
	}

	if !token.Valid { //Token is invalid, maybe not signed on this server
		if withRespond {
			u.Respond(w, u.Message("Token is not valid."), 403)
		}
		return false
	}

	// JWT

	if ok := AutorizationAPI(w, permissionType, tk, withRespond); !ok {
		return false
	}

	return true

}

// GetClaims return tk claims to refresh token
func GetClaims() *models.Claims {
	return tk
}
