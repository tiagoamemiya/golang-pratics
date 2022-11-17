package config

import (
	"gin-api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

const host = "localhost"
const user = "root"
const pass = "root"
const port = "5432"
const db = "root"

func DbConnect() {
	dsn := "host=" + host + " user=" + user + " password=" + pass + " dbname=" + db + " port=" + port + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Error connection to db: ", err.Error())
	}

	DB.AutoMigrate(models.Student{})
}
