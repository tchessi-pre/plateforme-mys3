package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/minio/minio-go/v7"
)

// UploadFile est le gestionnaire pour uploader un fichier dans MinIO et dans un répertoire local
func UploadFile(w http.ResponseWriter, r *http.Request) {
	bucket := r.URL.Query().Get("bucket")
	objectName := r.URL.Query().Get("object")
	if bucket == "" || objectName == "" {
		http.Error(w, "Bucket and object name are required", http.StatusBadRequest)
		return
	}

	// Initialiser le client MinIO
	minioClient, err := loadMinioClient()
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to initialize MinIO client: %v", err), http.StatusInternalServerError)
		return
	}

	// Récupérer le fichier téléversé depuis la requête
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 1. Créer un fichier temporaire pour sauvegarder le contenu du fichier uploadé
	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create temp file: %v", err), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Copier le fichier uploadé dans le fichier temporaire
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to copy file to temp file: %v", err), http.StatusInternalServerError)
		return
	}

	// Réouvrir le fichier temporaire pour MinIO (car le curseur a avancé)
	tempFile.Seek(0, 0)

	// 2. Upload du fichier vers MinIO
	_, err = minioClient.PutObject(context.Background(), bucket, objectName, tempFile, -1, minio.PutObjectOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to upload file to MinIO: %v", err), http.StatusInternalServerError)
		return
	}

	// 3. Enregistrer le fichier dans un répertoire local existant

	// Utiliser `os.Getwd` pour obtenir le répertoire de travail courant
	workingDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Unable to determine working directory", http.StatusInternalServerError)
		return
	}

	// Créer le chemin absolu vers le répertoire du bucket local (supposé exister)
	storageDir := filepath.Join(workingDir, "app", "storage", bucket)

	// Vérifier si le répertoire du bucket existe
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("Bucket directory %s does not exist", storageDir), http.StatusBadRequest)
		return
	}

	// Créer le chemin du fichier à l'intérieur du répertoire du bucket
	localFilePath := filepath.Join(storageDir, header.Filename)

	// Créer le fichier dans le répertoire existant
	localFile, err := os.Create(localFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create local file: %v", err), http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	// Réouvrir le fichier temporaire pour l'écriture locale
	tempFile.Seek(0, 0)

	// Copier le contenu du fichier temporaire dans le fichier local
	_, err = io.Copy(localFile, tempFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save file locally: %v", err), http.StatusInternalServerError)
		return
	}

	// Répondre avec succès
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File uploaded successfully to MinIO and local storage"))
}
