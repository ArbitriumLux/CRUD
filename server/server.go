package server

import (
	"github.com/jinzhu/gorm"
)

// Server settings

const PostgresConnectionParameters = "host=localhost port=5432 user=swkkd dbname=mydb sslmode=disable password=root"

var Db *gorm.DB
var err error

func DataBaseConnection() {
	//Database connection
	Db, err = gorm.Open("postgres", PostgresConnectionParameters)
	if err != nil {
		panic(err)
	}
}
