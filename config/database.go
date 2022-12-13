package config

import (
	"github.com/jutionck/code-run-2022/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	err = db.Debug().AutoMigrate(&model.BlogSpot{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
