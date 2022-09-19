package configuration

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func dbConn() (db *gorm.DB, err error) {
	return gorm.Open(postgres.New(postgres.Config{

		DSN:                  "host=localhost dbname=postgres port=5432 sslmode=disable TimeZone=Africa/Johannesburg",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "five_thirty.",
			SingularTable: true,
		},
	})
}

var Pdatabase *gorm.DB

func Connect() {
	db, err := dbConn()
	if err != nil {
		log.Fatal(err)
	}
	Pdatabase = db
}
