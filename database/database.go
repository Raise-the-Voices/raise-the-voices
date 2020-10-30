package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	// Gorm Postgres dialects
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("db_host"), os.Getenv("db_user"), os.Getenv("db_name"), os.Getenv("db_pass"))

	conn, err := gorm.Open(os.Getenv("db_type"), dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

}

// GetDB gorm
func GetDB() *gorm.DB {
	return db
}
