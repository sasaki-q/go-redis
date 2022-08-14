package database

import (
	"fmt"
	"os"

	"example.com/init/features/item"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		user,
		password,
		host,
		port,
		name,
	)

	db, err = gorm.Open("postgres", url)

	if err != nil {
		fmt.Print(err, url)
		panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(&item.ItemModel{})
}

func GetDB() *gorm.DB {
	return db
}
