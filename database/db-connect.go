package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func InItDB(){

	err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("postgresql://neondb_owner:%s@ep-gentle-morning-a5ivmvpt.us-east-2.aws.neon.tech/neondb?sslmode=require",password)

	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if(err!=nil){
		fmt.Println("Unable to connect database",err)
		os.Exit(1)
	}
	fmt.Println("database connected")
	DB = db
}