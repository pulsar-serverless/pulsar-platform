package postgres

import (
	"context"
	"pulsar/internal/core/domain/common"
	"pulsar/internal/core/domain/project"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectRepo struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) ProjectRepo {
	return ProjectRepo{db: db}
}

func Paginate[T any](pg *common.Pagination[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pg.PageNumber * pg.PageSize).Limit(pg.PageSize)
	}
}

func (repo *ProjectRepo) CreateProject(ctx context.Context, project *project.Project) error {
	result := repo.db.Create(project)
	return result.Error
}

func (repo *ProjectRepo) UpdateProject(ctx context.Context, projectId string, updatedProject *project.Project) (*project.Project, error) {
	var project project.Project
	result := repo.db.Where("id = ?", projectId).Updates(updatedProject)

	return &project, result.Error
}

func (repo *ProjectRepo) GetProject(ctx context.Context, projectId string) (*project.Project, error) {
	project := project.Project{ID: projectId}
	result := repo.db.Joins("SourceCode").First(&project)

	return &project, result.Error
}

func (repo *ProjectRepo) GetProjects(ctx context.Context, pageNumber int, pageSize int) (*common.Pagination[project.Project], error) {
	var projects []*project.Project

	pagination := &common.Pagination[project.Project]{
		PageSize:   pageSize,
		PageNumber: pageNumber,
	}

	// get total number of projects
	repo.db.Model(&project.Project{}).Count(&pagination.TotalPages)

	result := repo.db.Scopes(Paginate(pagination)).Find(&projects)
	pagination.Rows = projects

	return pagination, result.Error
}

func (repo *ProjectRepo) DeleteProject(ctx context.Context, projectId string) error {
	result := repo.db.Where("id = ?", projectId).Delete(&project.Project{})
	return result.Error
}

func (repo *ProjectRepo) UpdateSourceCode(ctx context.Context, id uuid.UUID, code *project.SourceCode) error {
	result := repo.db.Where("id = ?", id).Updates(code)
	return result.Error
}
