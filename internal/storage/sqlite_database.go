package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDatabase struct {
	db *gorm.DB
}

func (database *SqliteDatabase) InitDatabase(entities []interface{}) {
	db, err := gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	for _, model := range entities {
		err = db.AutoMigrate(model)
		if err != nil {
			panic("failed to migrate schema")
		}
	}
	database.db = db
}

func (database *SqliteDatabase) GetDatabase() *gorm.DB {
	return database.db
}
