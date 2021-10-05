package database

import (
	"filaments/config"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error

	port, err := strconv.ParseUint(config.Config("DB_PORT"), 10, 32)

	host := config.Config("DB_HOST")
	user := config.Config("DB_USER")
	password := config.Config("DB_PASSWORD")
	name := config.Config("DB_NAME")

	if err != nil {
		panic("failed to get port")
	}

	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}
