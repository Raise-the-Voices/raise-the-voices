package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Raise-the-Voices/raise-the-voices/database"
	"github.com/Raise-the-Voices/raise-the-voices/models"
	"github.com/Raise-the-Voices/raise-the-voices/routes"

	// Gorm Postgres dialects

	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	databaseTablesMigration()

	if database.GetDB().HasTable(&models.User{}) {
		seed()
	}

	// Routes
	routes := routes.Routes()

	// Server port reading
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "*", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "X-Requested-With", "Authorization"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	handler := c.Handler(routes)

	// Server listening
	err := http.ListenAndServe(":"+port, handler) //Launch the app.
	if err != nil {
		fmt.Print(err)
	}
}

// database tables migration
func databaseTablesMigration() {

	erroDB := database.GetDB().AutoMigrate(&models.Report{}, &models.Victim{}, &models.Incident{},
		&models.VictimMedia{}, &models.IncidentMedia{}, &models.VictimTranslation{}, &models.IncidentTranslation{},
		&models.User{}, &models.ResetPassword{}).Error

	if erroDB != nil {
		fmt.Println(erroDB)
	}
}

func seed() {
	userDefault := &models.User{
		Name:     "default",
		Password: "fX2Weey=y?*Z98&N",
		Email:    "exampleraisethevoices@gmail.com",
		Phone:    "11111111111",
		UserRole: "admin",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userDefault.Password), bcrypt.DefaultCost)
	userDefault.Password = string(hashedPassword)

	database.GetDB().Where(models.User{Email: "exampleraisethevoices@gmail.com"}).FirstOrCreate(&userDefault)

	if userDefault.ID <= 0 {
		fmt.Printf("Failed to create account, connection error.\n")
	}

	userDefault.Password = "" //delete password
	// fmt.Printf("Admin User created:\n%v ", userDefault)
}
