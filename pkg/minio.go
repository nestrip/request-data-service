package pkg

import (
	_ "embed"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"os"
)

var MinioClient *minio.Client
var ArchiveMinioClient *minio.Client

func ConnectToMinio() {
	var err error
	MinioClient, err = minio.New(os.Getenv("MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false, // since im connecting to minio via localhost, this is needed
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	ArchiveMinioClient, err = minio.New(os.Getenv("ARCHIVE_MINIO_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), ""),
		Secure: false, // since im connecting to minio via localhost, this is needed
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

//go:embed README.txt
var ReadMeFile string
