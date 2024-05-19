package postgres

import (
	"context"
	"math"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/project"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Database struct {
	conn *gorm.DB
}

func NewProjectRepo(db *gorm.DB) Database {
	return Database{conn: db}
}

func Paginate[T any](pg *common.Pagination[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((pg.PageNumber - 1) * pg.PageSize).Limit(pg.PageSize)
	}
}

func (repo *Database) CreateProject(ctx context.Context, project *project.Project) error {
	result := repo.conn.Create(project)
	return result.Error
}

func (repo *Database) UpdateProject(ctx context.Context, projectId string, updatedProject *project.Project) (*project.Project, error) {
	var project project.Project
	result := repo.conn.Where("id = ?", projectId).Updates(updatedProject)

	return &project, result.Error
}

func (repo *Database) UpdateProjectFields(ctx context.Context, projectId string, updatedProject map[string]interface{}) (*project.Project, error) {
	var project project.Project
	result := repo.conn.Model(&project).Where("id = ?", projectId).Updates(updatedProject)

	return &project, result.Error
}

func (repo *Database) GetProject(ctx context.Context, projectId string) (*project.Project, error) {
	project := project.Project{ID: projectId}
	result := repo.conn.Preload("SourceCode").Preload("EnvVariables").First(&project)

	return &project, result.Error
}

func (repo *Database) GetProjects(ctx context.Context, pageNumber int, pageSize int, userId string) (*common.Pagination[project.Project], error) {
	var projects []*project.Project

	pagination := &common.Pagination[project.Project]{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}

	// get total number of projects
	var count int64
	repo.conn.Model(&project.Project{}).Where(&project.Project{UserId: userId}).Count(&count)

	result := repo.conn.Scopes(Paginate(pagination)).
		Where(&project.Project{UserId: userId}).
		Order("updated_at desc").
		Find(&projects)

	pagination.Rows = projects
	pagination.TotalPages = int64(math.Ceil(float64(count) / float64(pageSize)))

	return pagination, result.Error
}

func (repo *Database) GetAllProjects(ctx context.Context, userId string) ([]*project.Project, error) {
	var projects []*project.Project

	result := repo.conn.
		Where(&project.Project{UserId: userId}).
		Find(&projects)

	return projects, result.Error
}

func (repo *Database) DeleteProject(ctx context.Context, projectId string) error {
	result := repo.conn.Where("id = ?", projectId).Delete(&project.Project{})
	return result.Error
}

func (repo *Database) UpdateSourceCode(ctx context.Context, id uuid.UUID, code *project.SourceCode) error {
	result := repo.conn.Where("id = ?", id).Updates(code)
	return result.Error
}
