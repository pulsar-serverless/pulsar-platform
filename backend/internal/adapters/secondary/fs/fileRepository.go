package fs

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"os"
	"path"
	"pulsar/internal/core/domain/project"
	"time"

	"github.com/docker/docker/pkg/archive"
	"github.com/mholt/archiver/v4"
	"github.com/otiai10/copy"
)

type ProjectFileRepository struct {
	rootPath           string
	dockerfileTemplate *template.Template
	starterCodePath    string
}

func NewProjectFileRepository(projectStoragePath, dockerfileTemplatePath, starterCodePath string) *ProjectFileRepository {
	template, err := template.ParseFiles(dockerfileTemplatePath)
	if err != nil {
		panic(fmt.Sprintf("Invalid docker config template. %v", err))
	}

	return &ProjectFileRepository{projectStoragePath, template, starterCodePath}
}

func (fileRepo *ProjectFileRepository) setupSourceFolder(project *project.Project) (string, error) {
	currentTimestamp := time.Now()
	sourcePath := path.Join(fileRepo.rootPath, currentTimestamp.Format(time.RFC3339)+project.Name)

	return sourcePath, os.Mkdir(sourcePath, os.ModePerm)
}

func (fileRepo *ProjectFileRepository) createDockerConfig(projectPath string, project *project.Project) error {

	dockerFile := path.Join(projectPath, "dockerfile")
	var file *os.File

	if _, err := os.Stat(dockerFile); os.IsNotExist(err) {
		if file, err = os.Create(dockerFile); err != nil {
			return err
		}
	} else {
		if file, err = os.OpenFile(dockerFile, os.O_WRONLY|os.O_TRUNC, 0644); err != nil {
			return err
		}
	}

	defer file.Close()

	return fileRepo.dockerfileTemplate.Execute(file, project)
}

func (fileRepo *ProjectFileRepository) SetupDefaultProject(project *project.Project) (string, error) {
	projectDir, err := fileRepo.setupSourceFolder(project)
	if err != nil {
		return "", err
	}

	// TODO: exclude dot files
	return projectDir, copy.Copy(fileRepo.starterCodePath, projectDir)
}

func (fileRepo *ProjectFileRepository) SetupCustomProjectCode(ctx context.Context, project *project.Project, zipFile *multipart.FileHeader) (string, error) {
	projectDir, err := fileRepo.setupSourceFolder(project)
	if err != nil {
		return "", err
	}

	src, err := zipFile.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	return projectDir, unzipSourceCode(ctx, projectDir, src)
}

func (fileRepo *ProjectFileRepository) CreateBuildContext(project *project.Project) (io.Reader, error) {
	if err := fileRepo.createDockerConfig(project.SourceCode.URI, project); err != nil {
		return nil, err
	}
	return archive.TarWithOptions(project.SourceCode.URI, &archive.TarOptions{})
}

func (fileRepo *ProjectFileRepository) ZipSourceCode(sourceDir string) (*os.File, error) {
	files, err := archiver.FilesFromDisk(nil, map[string]string{
		sourceDir: "",
	})

	if err != nil {
		return nil, err
	}

	out, err := os.CreateTemp("", "*.zip")
	if err != nil {
		return nil, err
	}
	defer out.Close()

	format := archiver.CompressedArchive{
		Archival: archiver.Zip{},
	}

	err = format.Archive(context.Background(), out, files)
	return out, err
}

func (fileRepo *ProjectFileRepository) RemoveSourceCode(sourceDir string) error {
	return os.RemoveAll(sourceDir)
}
