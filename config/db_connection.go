package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func InitDB() *gorm.DB {
	var err error

	env := os.Getenv("ENV")
	dbPasswd := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	if env != "dev" {
		// overwrite conf file not to use .env instead store in secure way such as AWS SSM paramstore or consul key value store
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USERNAME"),
		dbPasswd,
		dbHost,
		os.Getenv("MYSQL_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	db, errOpen := gorm.Open("mysql", dsn)
	if errOpen != nil {
		log.Fatal(errOpen.Error())
	}
	db.SingularTable(true)

	err = db.DB().Ping()

	if err != nil {
		log.Panic(err)
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Successfully connected to MySQL")
	}
	db.DB().SetMaxIdleConns(60)
	if env == "dev" {
		db.LogMode(true)
	}

	return db
}
