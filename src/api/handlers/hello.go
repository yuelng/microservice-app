package handlers

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"base/model"
	"time"
)

func Hello(c *gin.Context) {
	//p := fmt.Println

	//jsonTime := models.JSONTime(time.Now())
	//p(jsonTime)
	db := model.DB
	jsonTime := models.JSONTime{Time: time.Now()}
	account := models.Account{Tel: "Lbxxxc44450", Password: "sssdfss", StartAt: jsonTime}
	db.NewRecord(account)
	db.Create(&account)

	//p(account)
	//p(time.Now().UTC())

	// var account models.Account

	// db.First(&account, 5) // find product with id 1

	// time := account.StartAt.Format(time.RFC3339Nano)

	// utils.RenderJson(&account)

	// c.JSON(200, gin.H{"user": "hello"})
	//
	c.JSON(http.StatusOK, &account)
}
