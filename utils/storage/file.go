package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileStorage struct {
	Dir     string
	BaseUrl string
}

func (fs *FileStorage) Url(path string) (string, error) {
	if !fs.isFileExist(path) {
		return "", fmt.Errorf("file '%s' does not exist", path)
	}

	return filepath.Join(fs.BaseUrl, path), nil
}

func (fs *FileStorage) Path(path string) (string, error) {
	if !fs.isFileExist(path) {
		return "", fmt.Errorf("file '%s' does not exist", path)
	}

	return filepath.Join(fs.Dir, path), nil
}

func (fs *FileStorage) Read(path string) (*os.File, error) {
	if !fs.isFileExist(path) {
		return nil, fmt.Errorf("file '%s' does not exist", path)
	}

	return os.Open(filepath.Join(fs.Dir, path))
}

func (fs *FileStorage) Write(d io.Reader, path string) error {
	if fs.isFileExist(path) {
		return fmt.Errorf("file '%s' already exists", path)
	}

	var b []byte
	if _, err := d.Read(b); err != nil {
		return err
	}

	absPath := filepath.Join(fs.Dir, path)
	file, err := os.Create(absPath)
	if err != nil {
		return err
	}

	_, err = file.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileStorage) isFileExist(path string) bool {
	absPath := filepath.Join(fs.Dir, path)
	_, err := os.Stat(absPath)
	return os.IsExist(err)
}
