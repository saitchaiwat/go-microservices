package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"user-service/config"
	"user-service/database"
	"user-service/router"
)

func main() {
	config.LoadEnvironment()

	db := database.ConnectDatabase()
	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.Role{})

	defer database.DisconnectDatabase(db)

	app := fiber.New()

	router.SetupRoutes(app, db)

	port := viper.GetString("SERVER_PORT")
	app.Listen(port)
}
