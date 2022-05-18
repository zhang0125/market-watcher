package util

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var db *gorm.DB
	var err error
	dsn := conf.Dsn
	for {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("fail to connect database, err: " + err.Error())
			time.Sleep(time.Second * 3)
			continue
		}
		break
	}

	fmt.Println("[SUCCESS] DB connected")
	DB = db
}

func CloseDB() {
	fmt.Println("[SUCCESS] DB Closed")
}
