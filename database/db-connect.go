package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InItDB(){
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5440 sslmode=disable TimeZone=Asia/Shanghai"
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if(err!=nil){
		fmt.Println("Unable to connect database",err)
		os.Exit(1)
	}
	fmt.Println("database connected")
	DB = db
}