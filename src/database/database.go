package database

import (
	"cse-foodcourt/src/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func createDSN() string {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("failed to load env file")
	}
	// Load env properties
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// mysql", "root:password@tcp(localhost:3306)/testdb
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	return dsn
}

func createSchema(db *gorm.DB) {
	errDB := db.AutoMigrate(
		&model.Customer{},
		&model.Orders{},
		&model.PaymentDetails{},
		&model.Menu{},
		&model.Payment{},
		&model.Payment{},
		&model.Restaurant{})
	if errDB != nil {
		panic("Database can't create models")
	}
}

func DatabaseConection() {

	// crete dsn to database
	dsn := createDSN()

	// open database
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	// create schema
	createSchema(db)

	DB = db
	fmt.Println("Database Connected")

}
