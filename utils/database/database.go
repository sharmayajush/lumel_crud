package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/sharmayajush/lumel_crud/src/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetInstance() *gorm.DB {
	once.Do(func() {
		db = setupDatabase()
	})
	return db
}

func InitDBModels() {
	db.AutoMigrate(&model.Customer{}, model.Product{}, &model.Order{})
}

func setupDatabase() *gorm.DB {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	dbname := viper.GetString("database.dbname")
	searchpath := viper.GetString("database.schema")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC search_path=%s", host, username, password, dbname, port, searchpath)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Panicf("Error connecting to the postgressql database at %s:%s/%s", host, port, dbname)
	}

	return db
}
