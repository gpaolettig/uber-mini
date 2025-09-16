package main

import (
	"fmt"
	"os"
	"uber-mini/internal/adapters/http"
	"uber-mini/internal/app"
	"uber-mini/internal/core/driver"
	"uber-mini/internal/core/ride"
	"uber-mini/internal/core/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
		panic(err)
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	usr := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	fmt.Println("DB_HOST:", host)
	fmt.Println("DB_PORT:", port)
	fmt.Println("DB_USER:", usr)
	fmt.Println("DB_NAME:", dbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, usr, password, dbName, port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Failed to connect db %v", err))
	}
	if err := dbConn.AutoMigrate(&user.User{}, &driver.Driver{}, &driver.Car{}, &ride.Ride{}); err != nil {
		panic(fmt.Errorf("Failed to migrate db: %v", err))
	}

	driverHandler := app.Init(dbConn)
	r := gin.Default()
	http.InitRoutes(r, driverHandler)
	if err := r.Run(":8080"); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}

}
