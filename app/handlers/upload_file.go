package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/minio/minio-go/v7"
)

// UploadFile est le gestionnaire pour uploader un fichier dans MinIO
func UploadFile(w http.ResponseWriter, r *http.Request) {
	bucket := r.URL.Query().Get("bucket")
	objectName := r.URL.Query().Get("object")
	if bucket == "" || objectName == "" {
		http.Error(w, "Bucket and object name are required", http.StatusBadRequest)
		return
	}

	minioClient, err := loadMinioClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to initialize MinIO client: %v", err), http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload du fichier
	_, err = minioClient.PutObject(context.Background(), bucket, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to upload file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully"))
}
