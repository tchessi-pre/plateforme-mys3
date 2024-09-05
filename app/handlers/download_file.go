package handlers

import (
    "context"
    "fmt"
    "io"
    "net/http"

    "github.com/minio/minio-go/v7"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
    bucket := r.URL.Query().Get("bucket")
    key := r.URL.Query().Get("key")
    if bucket == "" || key == "" {
        http.Error(w, "Bucket and key are required", http.StatusBadRequest)
        return
    }

    // Initialiser le client MinIO
    minioClient, err := loadMinioClient()
    if err != nil {
        http.Error(w, fmt.Sprintf("Unable to initialize MinIO client: %v", err), http.StatusInternalServerError)
        return
    }

    // Obtenir l'objet depuis le bucket
    obj, err := minioClient.GetObject(context.Background(), bucket, key, minio.GetObjectOptions{})
    if err != nil {
        http.Error(w, fmt.Sprintf("Unable to download file: %v", err), http.StatusInternalServerError)
        return
    }
    defer obj.Close()

    // Définir l'en-tête pour le téléchargement du fichier
    w.Header().Set("Content-Disposition", "attachment; filename="+key)
    io.Copy(w, obj)
}
