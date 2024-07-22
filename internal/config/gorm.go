package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDatabase(config *viper.Viper) *gorm.DB {
	username := config.GetString("database.username")
	password := config.GetString("database.password")
	host := config.GetString("database.host")
	port := config.GetInt("database.port")
	database := config.GetString("database.name")
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable&lock_timeout=5000", username, password, host, port, database)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
