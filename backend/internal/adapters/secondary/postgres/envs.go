package postgres

import (
	"context"
	"pulsar/internal/core/domain/project"

	"gorm.io/gorm"
)

func (db *Database) OverwriteEnvVariables(ctx context.Context, projectId string, variables []*project.EnvVariable) ([]*project.EnvVariable, error) {
	err := db.conn.Transaction(func(db *gorm.DB) error {
		result := db.Where(&project.EnvVariable{ProjectID: projectId}).Delete(&project.EnvVariable{})
		if result.Error != nil {
			return result.Error
		}

		result = db.Create(&variables)
		return result.Error
	})

	if err != nil {
		return nil, err
	}

	return db.GetEnvVariables(ctx, projectId)
}

func (db *Database) GetEnvVariables(ctx context.Context, projectId string) ([]*project.EnvVariable, error) {
	var variables []*project.EnvVariable
	result := db.conn.Where(&project.EnvVariable{ProjectID: projectId}).Find(&variables)
	return variables, result.Error
}
