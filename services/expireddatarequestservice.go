package services

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/nestrip/request-data-service/db"
	"github.com/nestrip/request-data-service/ent/datarequest"
	"github.com/nestrip/request-data-service/pkg"
	"time"
)

func ExpiredDataRequestService(context context.Context) {
	for _, dR := range db.Client.DataRequest.Query().Where(datarequest.CreatedAtLTE(time.Now().AddDate(0, 0, -30))).AllX(context) {
		dR.Update().SetExpired(true).SaveX(context)
		err := pkg.MinioClient.RemoveObject(context, "data-requests", dR.CdnName, minio.RemoveObjectOptions{})
		if err != nil {
			fmt.Println(err)
		}
	}
}
