package database

import (
	"farm_management/entities"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect() {
	connString := fmt.Sprintf("server=%s;user=%s;password=%s;database=%s;", "host.docker.internal,1433", "sa", "sqlPwd@12#", "farms")
	Instance, err = gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")

	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	Instance.AutoMigrate(&entities.Users{})
	log.Println("Database Migration Completed...")
}
