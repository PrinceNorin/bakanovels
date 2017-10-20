package uploader

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime"
	"os"

	"github.com/PrinceNorin/bakanovels/utils/storage"
)

type ImageUploader struct {
	storage storage.Storage
}

func (iu *ImageUploader) Upload(base64Str string, u Uploadable) error {
	mtype, content, err := normalizeBase64(base64Str)
	if err != nil {
		return err
	}

	if !okMimeType(mtype) {
		return fmt.Errorf("invalid mime type '%s'", mtype)
	}

	unbased, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return err
	}

	exts, err := mime.ExtensionsByType(mtype)
	if err != nil {
		return err
	}

	fmt.Println(unbased)
	r := bytes.NewReader(unbased)
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	// w := &bytes.Buffer{}
	f, _ := os.Create(fmt.Sprintf("test%s", exts[0]))
	if mtype == "image/png" {
		png.Encode(f, img)
	} else {
		jpeg.Encode(f, img, nil)
	}

	return iu.storage.Write(r, fmt.Sprintf("%s%s", u.SavePath(), exts[0]))
}

func okMimeType(mtype string) bool {
	return mtype == "image/jpg" || mtype == "image/jpeg" || mtype == "image/png"
}
