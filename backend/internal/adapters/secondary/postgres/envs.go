package postgres

import (
	"context"
	"pulsar/internal/core/domain/project"

	"gorm.io/gorm"
)

func (db *Database) OverwriteEnvVariables(ctx context.Context, projectId string, variables []*project.EnvVariable) error {
	err := db.conn.Transaction(func(db *gorm.DB) error {
		result := db.Where(&project.EnvVariable{ProjectID: projectId}).Delete(&project.EnvVariable{})
		if result.Error != nil {
			return result.Error
		}

		result = db.Create(&variables)
		return result.Error
	})

	return err
}

func (db *Database) GetEnvVariables(ctx context.Context, variables []*project.EnvVariable) ([]*project.EnvVariable, error) {
	result := db.conn.Create(variables)
	return variables, result.Error
}
