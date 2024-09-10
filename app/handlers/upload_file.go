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

	// Vérifiez que le bucket et le nom de l'objet sont fournis
	if bucket == "" || objectName == "" {
		http.Error(w, "Bucket and object name are required", http.StatusBadRequest)
		return
	}

	// Vérifier que l'objectName ne se termine pas par un slash
	if len(objectName) > 0 && objectName[len(objectName)-1] == '/' {
		http.Error(w, "objectName cannot end with a '/'", http.StatusBadRequest)
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

	// Obtenir la taille du fichier
	fileSize := header.Size

	// Créer un fichier temporaire pour sauvegarder le contenu du fichier uploadé
	tempFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create temp file: %v", err), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name()) // Nettoyer le fichier temporaire après usage

	// Copier le fichier uploadé dans le fichier temporaire
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to copy file to temp file: %v", err), http.StatusInternalServerError)
		return
	}

	// Réouvrir le fichier temporaire pour MinIO (car le curseur a avancé)
	tempFile.Seek(0, 0)

	// Upload du fichier vers MinIO
	_, err = minioClient.PutObject(context.Background(), bucket, objectName, tempFile, fileSize, minio.PutObjectOptions{})
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to upload file to MinIO: %v", err), http.StatusInternalServerError)
		return
	}

	// Enregistrer le fichier dans un répertoire local existant
	workingDir, err := os.Getwd()
	if err != nil {
		http.Error(w, "Unable to determine working directory", http.StatusInternalServerError)
		return
	}

	// Créer le chemin absolu vers le répertoire du bucket local
	storageDir := filepath.Join(workingDir, "app", "storage", bucket)

	// Créer le répertoire si nécessaire
	if _, err := os.Stat(storageDir); os.IsNotExist(err) {
		err = os.MkdirAll(storageDir, os.ModePerm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create bucket directory %s: %v", storageDir, err), http.StatusInternalServerError)
			return
		}
	}

	// Utiliser `filepath.Base` pour s'assurer que seul le nom du fichier est utilisé
	localFilePath := filepath.Join(storageDir, filepath.Base(header.Filename))

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
