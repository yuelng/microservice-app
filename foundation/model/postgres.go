package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

const dialect = "postgres"

var DB *gorm.DB
var once sync.Once

// use global connection
// or pass db from handler to models, by struct method or args
// 对于使用:=定义的变量，如果新变量p与那个同名已定义变量 (这里就是那个全局变量p)不在一个作用域中时，那么golang会新定义这个变量p，遮盖住全局变量p

func InitDB(url string) {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(dialect, url)
		if err != nil {
			panic(fmt.Errorf("fatal error when connecting database: %s", err))
		}

		//db.LogMode(true)
		DB.DB().SetMaxOpenConns(100)
		DB.DB().SetMaxIdleConns(10)

	})
}
