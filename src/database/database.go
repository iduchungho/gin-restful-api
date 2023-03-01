package database

import (
	"cse-foodcourt/src/model"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}
var databaseConnection *gorm.DB

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
		&model.Restaurant{},
		&model.Comment{})
	if errDB != nil {
		panic("Database can't create models")
	}
}

func createConnect() *gorm.DB {
	dsn := createDSN()
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	createSchema(db)
	return db
}

/*
GetConnection
Design Pattern: Singleton which ensures that
only one object of its kind exists and provides a single point
of access to it for any other code.
*/
func GetConnection() *gorm.DB {
	if databaseConnection == nil {
		lock.Lock()
		// Ensure that the instance hasn't yet been
		// initialized by another thread while this one
		// has been waiting for the lock's release.
		defer lock.Unlock()
		if databaseConnection == nil {
			databaseConnection = createConnect()
			fmt.Println("Database Connected")

		} else {
			return databaseConnection
		}
	}
	return databaseConnection
}
