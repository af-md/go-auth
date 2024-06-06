package services

import (
	"errors"
	"fmt"
	"log"

	"go-auth/model"

	"github.com/caarlos0/env/v11"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDB() (*gorm.DB, error) {

	dbCfg := model.DBConfig{}

	if err := env.Parse(&dbCfg); err != nil {
		err = errors.New("failed to parse db config: " + err.Error())
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.DbName)

	log.Printf("create dsn: %s", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		err = errors.New("failed to connect database: " + err.Error())
		return nil, err
	}
	return db, nil
}
