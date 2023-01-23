package zipReader

import (
	"archive/zip"
	"io"
)

const dataFileName = "xl/sharedStrings.xml"

var dataNotFoundErr error

type Reader struct {
	archive *zip.ReadCloser
}

func New() *Reader {
	return &Reader{}
}

func (r *Reader) UnzipFile(pathToExcelFile string) (io.ReadCloser, error) {
	archive, err := zip.OpenReader(pathToExcelFile)
	if err != nil {
		return nil, err
	}

	r.archive = archive

	for _, f := range archive.File {
		if f.Name == dataFileName {
			return f.Open()
		}
	}

	return nil, dataNotFoundErr
}

func (r *Reader) CloseArchive() {
	if r.archive != nil {
		r.archive.Close()
	}
}
