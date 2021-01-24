package routes

import (
	"net/http"
	"time"

	"github.com/Raise-the-Voices/raise-the-voices/controllers"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"

	"github.com/gorilla/mux"
)

// Routes all routes
func Routes() *mux.Router {
	router := mux.NewRouter()

	// API VERSION
	s := router.PathPrefix("/v1").Subrouter()

	// Allow at most 100 request per 1 second.
	limiter := tollbooth.NewLimiter(100, &limiter.ExpirableOptions{DefaultExpirationTTL: time.Second})

	// Set a custom content-type.
	limiter.SetMessageContentType("application/json; charset=utf-8")

	// -------------------Reports Routes---------------------------------------

	// Create a new Victim Report.
	s.Handle("/reports", tollbooth.LimitFuncHandler(limiter, controllers.CreateVictimReport)).Methods("POST")

	// Modify the report state of a victim report.
	s.Handle("/reports/{idreport}", tollbooth.LimitFuncHandler(limiter, controllers.ModifyReportState)).Methods("PATCH")

	// Update data from a specific report.
	s.Handle("/reports/{idreport}", tollbooth.LimitFuncHandler(limiter, controllers.UpdateReportData)).Methods("PUT")

	// Get reports depending on query.
	s.Handle("/reports", tollbooth.LimitFuncHandler(limiter, controllers.GetReport)).Methods("GET")

	// Delete a victim report by id
	s.Handle("/reports/{idreport}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteReport)).Methods("DELETE")

	// -------------------Victims Routes---------------------------------------

	// update data from a specific victim.
	s.Handle("/victims/{idvictim}", tollbooth.LimitFuncHandler(limiter, controllers.UpdateVictimData)).Methods("PUT")

	// Get victims depending on query.
	s.Handle("/victims", tollbooth.LimitFuncHandler(limiter, controllers.GetVictims)).Methods("GET")

	// Delete a victim by her id
	s.Handle("/victims/{idvictim}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteVictim)).Methods("DELETE")

	// Download pdf file with information on one or several victims.
	s.Handle("/victims/pdf", tollbooth.LimitFuncHandler(limiter, controllers.GetVictimsPDF)).Methods("GET")

	// Upload profile photo of victim according to the victim's ID. just should allowed PNG or JPG format file.
	s.Handle("/victims/profile-img/{idvictim}", tollbooth.LimitFuncHandler(limiter, controllers.UploadProfilePhoto)).Methods("POST")

	// Delete the victim's profile photo according to the victim's identification.
	s.Handle("/victims/profile-img/{idvictim}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteProfilePhoto)).Methods("DELETE")

	// -------------------Victim_Translations Routes---------------------------------------

	// Create a new Victim_Translation.
	s.Handle("/victims/{idvictim}/victim-translations", tollbooth.LimitFuncHandler(limiter, controllers.AddVictimTranslationController)).Methods("POST")

	// Update data from a specific translation.
	s.Handle("/victim-translations/{idvictim-translation}", tollbooth.LimitFuncHandler(limiter, controllers.UpdateTranslationData)).Methods("PUT")

	// Get victim_translations depending on query.
	s.Handle("/victim-translations", tollbooth.LimitFuncHandler(limiter, controllers.GetVictimTranslations)).Methods("GET")

	// Delete a victim_translation by its id
	s.Handle("/victim-translations/{idvictim-translation}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteTranslation)).Methods("DELETE")

	// -------------------Victim_Media Routes---------------------------------------

	// Add a new Media of a Victim.
	s.Handle("/victims/{idvictim}/victimmedias", tollbooth.LimitFuncHandler(limiter, controllers.AddVictimMediaController)).Methods("POST")

	// Delete a victim Media by its id
	s.Handle("/victimmedias/{idvictimmedia}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteVictimMediaController)).Methods("DELETE")

	// Get victim_Media depending on query.
	s.Handle("/victimmedias", tollbooth.LimitFuncHandler(limiter, controllers.GetVictimMediaController)).Methods("GET")

	// Upload Victim Media File.
	s.Handle("/victimmedias/upload", tollbooth.LimitFuncHandler(limiter, controllers.UploadVictimFile)).Methods("POST")

	// -------------------Incidents Routes---------------------------------------

	// Add a new Incident of a Victim.
	s.Handle("/victims/{idvictim}/incidents", tollbooth.LimitFuncHandler(limiter, controllers.AddIncidentController)).Methods("POST")

	// Update data from a specific incident.
	s.Handle("/incidents/{idincident}", tollbooth.LimitFuncHandler(limiter, controllers.UpdateIncidentData)).Methods("PUT")

	// Get incidents depending on query.
	s.Handle("/incidents", tollbooth.LimitFuncHandler(limiter, controllers.GetIncidents)).Methods("GET")

	// Delete a incident by its id
	s.Handle("/incidents/{idincident}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteIncident)).Methods("DELETE")

	// -------------------Incident_Translations Routes---------------------------------------

	// Add a new Incident_Translation.
	s.Handle("/incidents/{idincident}/incident-translations", tollbooth.LimitFuncHandler(limiter, controllers.AddIncidentTranslationController)).Methods("POST")

	// Update data from a specific translation of a incident.
	s.Handle("/incident-translations/{idincident-translation}", tollbooth.LimitFuncHandler(limiter, controllers.UpdateIncidentTranslation)).Methods("PUT")

	// Get incidents_translations depending on query.
	s.Handle("/incident-translations", tollbooth.LimitFuncHandler(limiter, controllers.GetIncidentTranslations)).Methods("GET")

	// Delete a incident_translation by its id
	s.Handle("/incident-translations/{idincident-translation}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteIncidentTranslation)).Methods("DELETE")

	// -------------------Incident_Medias Routes---------------------------------------

	// Add a new Media of a Incident.
	s.Handle("/incidents/{idincident}/incident-medias", tollbooth.LimitFuncHandler(limiter, controllers.AddIncidentMedia)).Methods("POST")

	// Delete a incident Media by its id
	s.Handle("/incident-medias/{idincident-media}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteIncidentMedia)).Methods("DELETE")

	// Get incident_Media depending on query.
	s.Handle("/incident-medias", tollbooth.LimitFuncHandler(limiter, controllers.GetIncidentMedia)).Methods("GET")

	// Upload Incident Media File.
	s.Handle("/incident-medias/upload", tollbooth.LimitFuncHandler(limiter, controllers.UploadIncidentFile)).Methods("POST")

	// -------------------User Routes---------------------------------------

	// Signin User.
	s.Handle("/signin", tollbooth.LimitFuncHandler(limiter, controllers.Authenticate)).Methods("POST")

	// Refresh Token.
	s.Handle("/refresh", tollbooth.LimitFuncHandler(limiter, controllers.RefreshToken)).Methods("POST")

	// Create a new user.
	s.Handle("/users", tollbooth.LimitFuncHandler(limiter, controllers.CreateUser)).Methods("POST")

	// Modify password from a user.
	s.Handle("/users/change-password", tollbooth.LimitFuncHandler(limiter, controllers.UpdateUserPassword)).Methods("POST")

	// Modify data of a user.
	s.Handle("/users/{iduser}", tollbooth.LimitFuncHandler(limiter, controllers.UpdateUserData)).Methods("PUT")

	// Get users.
	s.Handle("/users", tollbooth.LimitFuncHandler(limiter, controllers.GetUsers)).Methods("GET")

	// Delete a user
	s.Handle("/users/{iduser}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteUser)).Methods("DELETE")

	// Get user email to reset password
	s.Handle("/password/email", tollbooth.LimitFuncHandler(limiter, controllers.PasswordEmail)).Methods("POST")

	// Get user token and response if is valid or not
	s.Handle("/password/reset/{token}", tollbooth.LimitFuncHandler(limiter, controllers.ResetToken)).Methods("GET")

	// Reset Password from a user
	s.Handle("/password/reset", tollbooth.LimitFuncHandler(limiter, controllers.ResetPassword)).Methods("POST")

	// -------------------Options Routes-----------------------------

	// Add a new option.
	s.Handle("/options", tollbooth.LimitFuncHandler(limiter, controllers.AddOptionController)).Methods("POST")

	// Delete an option.
	s.Handle("/options/{optionid}", tollbooth.LimitFuncHandler(limiter, controllers.DeleteOptionController)).Methods("DELETE")

	// Get all options.
	s.Handle("/options", tollbooth.LimitFuncHandler(limiter, controllers.GetAllOptionsController)).Methods("GET")

	// -------------------Error Routes---------------------------------------

	// if there was not a router found this trigger this error route.
	s.NotFoundHandler = http.HandlerFunc(controllers.NotFound)

	return router
}
