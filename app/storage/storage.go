package storage

import (
	"context"
	"fmt"
	"log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func init() {
	var err error
	minioClient, err = minio.New("http://minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("admin", "admin123", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

// CreateBucket creates a new bucket in MinIO.
func CreateBucket(bucketName string) error {
	ctx := context.Background()
	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			// Bucket already exists
			return nil
		}
		return err
	}
	fmt.Printf("Successfully created %s\n", bucketName)
	return nil
}

// Other functions for UploadFile, ListFiles, DownloadFile, DeleteFile...
