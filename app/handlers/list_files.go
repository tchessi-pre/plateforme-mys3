package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/minio/minio-go/v7"
)

// ListFiles est le gestionnaire pour lister les fichiers dans un bucket
func ListFiles(w http.ResponseWriter, r *http.Request) {
	bucket := r.URL.Query().Get("bucket")
	if bucket == "" {
		http.Error(w, "Bucket name is required", http.StatusBadRequest)
		return
	}

	minioClient, err := loadMinioClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to initialize MinIO client: %v", err), http.StatusInternalServerError)
		return
	}

	// Obtenir la liste des objets dans le bucket
	objectCh := minioClient.ListObjects(context.Background(), bucket, minio.ListObjectsOptions{})

	objects := []string{}
	for object := range objectCh {
		if object.Err != nil {
			http.Error(w, fmt.Sprintf("Error listing objects: %v", object.Err), http.StatusInternalServerError)
			return
		}
		objects = append(objects, object.Key)
	}

	// Retourner la liste des objets en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objects)
}
