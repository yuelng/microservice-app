package models

import (
	"base/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func InitDB() {
	url := "postgresql://go_web_dev:go_web_dev@localhost/go_web_dev?sslmode=disable"

	model.InitDB(url)
	InitSchema()
	Seed()
}

func InitSchema() {
	db := model.DB
	db.LogMode(true)
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis;")
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis_topology;")

	db.AutoMigrate(&Account{}, &Location{})
}

func Seed() {
	InitAccount()
}
