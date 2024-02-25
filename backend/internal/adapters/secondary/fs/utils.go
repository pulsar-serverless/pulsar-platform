package fs

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/archiver/v4"
)

func unzipSourceCode(ctx context.Context, base string, src multipart.File) error {
	format := archiver.Zip{}
	handler := func(ctx context.Context, f archiver.File) error {
		targetPath := filepath.Join(base, f.NameInArchive)

		if strings.HasPrefix(f.NameInArchive, "__MACOSX/") {
			return nil
		}

		if f.FileInfo.IsDir() {
			os.MkdirAll(targetPath, f.Mode())
		} else {
			inputFile, err := f.Open()
			if err != nil {
				return err
			}

			outputFile, err := os.OpenFile(
				targetPath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				f.Mode(),
			)

			if err != nil {
				return err
			}

			defer outputFile.Close()

			_, err = io.Copy(outputFile, inputFile)
			if err != nil {
				return err
			}
		}

		return nil
	}

	return format.Extract(ctx, src, nil, handler)
}
