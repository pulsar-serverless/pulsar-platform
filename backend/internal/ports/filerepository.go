package ports

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"pulsar/internal/core/domain/project"
)

type IFileRepository interface {
	SetupDefaultProject(project *project.Project) (string, error)
	SetupCustomProjectCode(ctx context.Context, project *project.Project, zipFile *multipart.FileHeader) (string, error)
	CreateBuildContext(project *project.Project) (io.Reader, error)
	ZipSourceCode(sourceDir string) (*os.File, error)
	RemoveSourceCode(sourceDir string) error
}
