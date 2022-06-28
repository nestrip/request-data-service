package services

import (
	"archive/zip"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/nestrip/request-data-service/db"
	"github.com/nestrip/request-data-service/ent/user"
	"github.com/nestrip/request-data-service/pkg"
)

func DataRequestService(context context.Context) {
	users := db.Client.User.Query().Where(user.RequestingData(true)).
		WithDomains().
		WithUsableDomains().
		WithMotds().
		WithTestimonial().
		AllX(context)

	for _, u := range users {
		fmt.Println("Processing data request for user", u.Username)

		buf := new(bytes.Buffer)
		zipFile := zip.NewWriter(buf)
		_ = pkg.AddFileToZip(zipFile, "README.txt", []byte(pkg.ReadMeFile))

		accountBytes, _ := json.MarshalIndent(u, "", "  ")
		_ = pkg.AddFileToZip(zipFile, "account.json", accountBytes)

		files := u.QueryFiles().AllX(context)
		fileBytes, _ := json.MarshalIndent(files, "", "  ")
		_ = pkg.AddFileToZip(zipFile, "files.json", fileBytes)

		for _, f := range files {
			fileData, err := pkg.MinioClient.GetObject(context, "uploads", f.CdnFileName, minio.GetObjectOptions{})
			if err == nil {
				_ = pkg.AddMinioFileToZip(zipFile, "uploads/"+f.CdnFileName, fileData)
			}
		}

		_ = zipFile.Close()

		fileName := uuid.New().String() + ".zip"
		u.Update().SetRequestingData(false).SaveX(context)
		dR := db.Client.DataRequest.Create().SetCreator(u).SetCdnName(fileName).SetExpired(false).SaveX(context)
		_, err := pkg.MinioClient.PutObject(context, "data-requests", fileName, buf, int64(buf.Len()), minio.PutObjectOptions{})

		if err != nil {
			panic(err)
		}

		pkg.SendDataRequestComplete(u, dR, context)
	}
}
