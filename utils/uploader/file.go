package uploader

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"mime"
	"strings"

	"github.com/PrinceNorin/bakanovels/utils/storage"
)

type FileUploader struct {
	storage storage.Storage
}

func (fu *FileUploader) Upload(base64Str string, u Uploadable) error {
	mtype, content, err := normalizeBase64(base64Str)
	if err != nil {
		return err
	}

	unbased, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	exts, err := mime.ExtensionsByType(mtype)
	if err != nil {
		return err
	}

	return fu.storage.Write(bytes.NewReader(unbased),
		fmt.Sprintf("%s%s", u.SavePath(), exts[0]))
}

func normalizeBase64(str string) (string, string, error) {
	p := strings.Split(str, ",")
	if len(p) < 2 {
		return "", "", errors.New("invalid base64 string")
	}

	pp := strings.Split(p[0], ";")
	if len(pp) < 2 {
		return "", "", errors.New("invalid base64 string")
	}

	ppp := strings.Split(pp[0], ":")
	if len(ppp) < 2 {
		return "", "", errors.New("invalid base64 string")
	}

	return ppp[1], p[1], nil
}
