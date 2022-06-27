package pkg

import (
	"archive/zip"
	"bytes"
	"github.com/minio/minio-go/v7"
	"io"
)

func CreateZipFile() *zip.Writer {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	return w
}

func AddFileToZip(w *zip.Writer, fileName string, fileBytes []byte) error {
	f, err := w.Create(fileName)
	if err != nil {
		return err
	}

	_, err = f.Write(fileBytes)
	if err != nil {
		return err
	}

	return nil
}

func AddMinioFileToZip(w *zip.Writer, fileName string, object *minio.Object) error {
	f, err := w.Create(fileName)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, object)
	if err != nil {
		return err
	}

	return nil
}
