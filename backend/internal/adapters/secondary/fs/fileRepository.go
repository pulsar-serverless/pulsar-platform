package fs

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"os"
	"path"
	"pulsar/internal/core/domain/billing"
	"pulsar/internal/core/domain/project"
	"time"

	"github.com/docker/docker/pkg/archive"
	"github.com/go-pdf/fpdf"
	"github.com/mholt/archiver/v4"
	"github.com/otiai10/copy"
)

type ProjectFileRepository struct {
	rootPath           string
	sitePath           string
	dockerfileTemplate *template.Template
	starterCodePath    string
	templateSitePath   string
	invoicePath        string
}

func NewProjectFileRepository(projectStoragePath, projectSitePath, dockerfileTemplatePath, starterCodePath, templateSitePath, invoiceStoragePath string) *ProjectFileRepository {
	template, err := template.ParseFiles(dockerfileTemplatePath)
	if err != nil {
		panic(fmt.Sprintf("Invalid docker config template. %v", err))
	}

	return &ProjectFileRepository{projectStoragePath, projectSitePath, template, starterCodePath, templateSitePath, invoiceStoragePath}
}

func (fileRepo *ProjectFileRepository) setupSourceFolder(project *project.Project) (string, error) {
	currentTimestamp := time.Now()
	sourcePath := path.Join(fileRepo.rootPath, currentTimestamp.Format(time.RFC3339)+project.Name)

	return sourcePath, os.Mkdir(sourcePath, os.ModePerm)
}

func (fileRepo *ProjectFileRepository) setupSiteFolder(project *project.Project) (string, error) {
	currentTimestamp := time.Now()
	sourcePath := path.Join(fileRepo.sitePath, currentTimestamp.Format(time.RFC3339)+project.Name)

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

func (fileRepo *ProjectFileRepository) SetupDefaultProjectSite(project *project.Project) (string, error) {
	siteDir, err := fileRepo.setupSiteFolder(project)
	if err != nil {
		return "", err
	}

	// TODO: exclude dot files
	return siteDir, copy.Copy(fileRepo.templateSitePath, siteDir)
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

func (fileRepo *ProjectFileRepository) SetupCustomSiteAssets(ctx context.Context, project *project.Project, zipFile *multipart.FileHeader) (string, error) {
	siteDir, err := fileRepo.setupSiteFolder(project)
	if err != nil {
		return "", err
	}

	assets, err := zipFile.Open()
	if err != nil {
		return "", err
	}
	defer assets.Close()

	return siteDir, unzipSourceCode(ctx, siteDir, assets)
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

func (fileRepo *ProjectFileRepository) SaveInvoicePDF(invoice *billing.Invoice, pdf *fpdf.Fpdf) (string, error) {
	filePath := path.Join(fileRepo.invoicePath, "invoice-"+invoice.UsageMonth+invoice.ProjectID+".pdf")

	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
