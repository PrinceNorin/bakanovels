package storage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/PrinceNorin/bakanovels/config"
)

type Storage interface {
	Url(id string) (string, error)
	Path(id string) (string, error)
	Read(id string) (*os.File, error)
	Write(d io.Reader, path string) error
}

var once sync.Once
var _storage Storage

func Get() Storage {
	once.Do(func() {
		// TODO: maybe switch to cloud base storage
		c := config.Get()
		root, _ := os.Getwd()
		_storage = &FileStorage{
			Dir:     filepath.Join(root, c.Storage.Dir),
			BaseUrl: fmt.Sprintf("http://%s:%d", c.Host, c.Port),
		}
	})

	return _storage
}
