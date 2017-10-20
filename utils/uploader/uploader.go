package uploader

import (
	"github.com/PrinceNorin/bakanovels/utils/storage"
)

type Uploadable interface {
	SavePath() string
	SaveColumnName() string
	SetValue(value string)
}

type Uploader interface {
	Upload(base64Str string, u Uploadable) error
}

func NewUploader() Uploader {
	return &FileUploader{
		storage: storage.Get(),
	}
}

func NewImageUploader() Uploader {
	return &ImageUploader{
		storage: storage.Get(),
	}
}
