package ports

import (
	"io"
	"mime/multipart"
	"pulsar/internal/core/domain/project"
)

type IFileRepository interface {
	InstallProject(project *project.Project, sourceFile *multipart.FileHeader) error
	InstallDefaultProject(project *project.Project) (string, error)
	CreateBuildContext(projectDir string) (io.Reader, error)
}
