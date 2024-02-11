package fs

import (
	"archive/zip"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func extractZippedProject(sourcePath string, targetPath string) error {
	zipReader, err := zip.OpenReader(sourcePath)
	if err != nil {
		return err
	}

	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		path, filePath := filepath.Split(file.Name)
		if hasHiddenFolder(path) || strings.HasPrefix(filePath, ".") {
			continue
		}

		extractedFilePath := filepath.Join(targetPath, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)

			if err != nil {
				return err
			}

			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func copyProject(sourceFile *multipart.FileHeader) (string, error) {
	tmpDir, err := os.MkdirTemp("", "project")
	if err != nil {
		return "", err
	}

	src, err := sourceFile.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	destination, err := os.Create(filepath.Join(tmpDir, sourceFile.Filename))
	if err != nil {
		return "", err
	}

	defer destination.Close()
	if io.Copy(destination, src); err != nil {
		return "", err
	}

	return destination.Name(), nil
}

func hasHiddenFolder(path string) bool {
	hiddenFolderRegex := regexp.MustCompile(`/\..*`)
	return hiddenFolderRegex.MatchString(path)
}
