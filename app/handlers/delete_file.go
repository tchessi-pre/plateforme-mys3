package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/minio/minio-go/v7"
)

// DeleteFile est le gestionnaire pour supprimer un fichier dans MinIO
func DeleteFile(w http.ResponseWriter, r *http.Request) {
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

	// Supprimer l'objet
	err = minioClient.RemoveObject(context.Background(), bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to delete file: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File deleted successfully"))
}
