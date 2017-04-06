package services

import (
	"api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
)

var once sync.Once

const dialect = "postgres"
const url = "postgresql://go_web_dev:go_web_dev@localhost/go_web_dev?sslmode=disable"

func InitDB() *gorm.DB {
	var DB *gorm.DB
	var err error

	once.Do(func() {
		DB, err = gorm.Open(dialect, url)
		if err != nil {
			panic(fmt.Errorf("fatal error when connecting database: %s", err))
		}

		DB.DB().SetMaxOpenConns(100)
		DB.DB().SetMaxIdleConns(10)
	})
	return DB
}

func InitSchema() {
	db := InitDB()
	db.LogMode(true)
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis;")
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis_topology;")

	db.AutoMigrate(&models.Account{}, &models.Location{})
}

func Seed() {
	models.InitAccount()
}
