package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func ConnectDatabase() *gorm.DB {
	dsn := viper.GetString("DATABASE_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	return db
}

func DisconnectDatabase(gorm *gorm.DB) {
	db, err := gorm.DB()
	if err != nil {
		log.Fatalf("Failed to disconnect: %v", err)
	}

	db.Close()
}
