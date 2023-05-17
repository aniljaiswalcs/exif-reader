package filereader

import (
	"os"
	"path/filepath"
)

// Reader defines exif reader reader
type Reader interface {
	FilePath(path string) ([]string, error)
}

// reader implements Reader interface.
type reader struct{}

// NewReader returns a new instance of reader.
func NewReader() Reader {
	return &reader{}
}

var filespath []string

func (r *reader) FilePath(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}

	for _, file := range files {
		if file.IsDir() {
			subDirectoryPath := filepath.Join(path, file.Name())
			r.FilePath(subDirectoryPath)
		} else {
			filepath := filepath.Join(path, file.Name())
			filespath = append(filespath, filepath)
		}
	}

	return filespath, nil
}
