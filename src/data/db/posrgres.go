package db

import (
	"clean-web-api/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	cnn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,	
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode,
	)
	dbClient , err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Postgres.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.Postgres.MaxOpenConnections)
	sqlDB.SetConnMaxIdleTime(cfg.Postgres.ConnMaxLifeTime * time.Minute)

	log.Println("Database Connection Successfully")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb(){
	conn , _ := dbClient.DB()
	conn.Close()
}