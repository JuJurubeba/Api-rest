package config

import (
	"os"

	"github.com/JuJurubeba/Api-rest/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("slite")
	dbPath := "./db/main.db"


	_, err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err 
		}
		file , err := os.Create(dbPath)
		if err != nil{
			return nil, err
		}
		file.Close()


	}

	//create database and conect

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil{
		logger.Error("sqlite opening error: &v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil{
		logger.Error("sqlite automigration error: &v", err)
		return nil, err
	}
	//return db
	return db, nil
}
