package models

import (
	_ "api/services"
)
type Account struct {
	Base
	Tel      string   `json:"tel" gorm:"not null; unique;"`
	Password string   `json:"-" gorm:"not null;"`
	StartAt  JSONTime `json:"start_at"`
}

func AddAccount()  {

}