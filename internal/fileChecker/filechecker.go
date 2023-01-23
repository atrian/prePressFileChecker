package fileChecker

import (
	"fmt"
	"os"
)

type Checker struct {
	FileExtension string
	ImageFolder   string
	FilesToCheck  []string
	FilesExist    []string
	FilesNotExist []string
}

func New(fileExtension string, imageFolder string) *Checker {
	ch := Checker{
		FileExtension: fileExtension,
		ImageFolder:   imageFolder,
	}

	return &ch
}

func (ch *Checker) Load(files []string) *Checker {
	ch.FilesToCheck = files
	return ch
}

func (ch *Checker) CheckFiles() *Checker {
	exist := make([]string, 0, len(ch.FilesToCheck))
	notExist := make([]string, 0, len(ch.FilesToCheck))

	for _, file := range ch.FilesToCheck {
		filePath := fmt.Sprintf("%v\\%v.%v", ch.ImageFolder, file, ch.FileExtension)
		if fileExists(filePath) {
			exist = append(exist, file)
		} else {
			notExist = append(notExist, file)
		}
	}

	ch.FilesExist = exist
	ch.FilesNotExist = notExist

	return ch
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
