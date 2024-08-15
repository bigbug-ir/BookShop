package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*****************************************************************/

type DbInstance struct {
	DB *gorm.DB
}

/*****************************************************************/

var Database DbInstance

/*****************************************************************/

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("MyApp.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: \n", err)
		os.Exit(2)
	}
	log.Println("Connecting database: ", db.Name())
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	log.Println("`Successfully connected to the database")
	Database = DbInstance{DB: db}
}

/*****************************************************************/
