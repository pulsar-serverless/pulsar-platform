package postgres

import (
	"pulsar/internal/core/domain/project"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to the database")
	}

	err = db.AutoMigrate(&project.Project{})
	if err != nil {
		panic("failed to perform migrations: " + err.Error())
	}

	return db
}
