package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	*gorm.DB
}

func InitDB() *DB {
	url := "postgresql://go_web_dev:go_web_dev@localhost/go_web_dev?sslmode=disable"

	dialect := "postgres"

	db, err := gorm.Open(dialect, url)
	if err != nil {
		panic(fmt.Errorf("fatal error when connecting database: %s", err))
	}

	// db.LogMode(true)
	return &DB{db}
}

func (db *DB) InitSchema() {
	db.LogMode(true)
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis;")
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis_topology;")
	db.AutoMigrate(&Account{}, &Location{})

}

func (db *DB) Seed() {
	InitAccount()
}
