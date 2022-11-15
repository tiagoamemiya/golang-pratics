package models

import (
	"api/database"
)

type Person struct {
	Id       int    `json:"id"`
	Nome     string `json:"name"`
	Historia string `json:"bio"`
}

const tableName = "personalidades"

func PersonsAll() []Person {
	database.DbConnect()
	persons := []Person{}

	database.DB.Table(tableName).Find(&persons)

	return persons
}

func PersonGet(id int) Person {
	database.DbConnect()
	var p Person
	database.DB.Table(tableName).First(&p, "id = ?", id)
	return p
}

func PersonNew(p *Person) {
	database.DbConnect()
	database.DB.Table(tableName).Create(&p)
}

func PersonDelete(p Person, id int) {
	database.DbConnect()
	database.DB.Table(tableName).Delete(p, id)
}

func PersonUpdate(p *Person) {
	database.DbConnect()
	database.DB.Table(tableName).Save(&p)
}
