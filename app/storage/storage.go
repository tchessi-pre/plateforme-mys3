package storage

import (
	"os"
	"path/filepath"
)

// CreateBucket crée un nouveau bucket (dossier) dans le système de fichiers.
func CreateBucket(bucketName string) error {
	return os.MkdirAll(filepath.Join("storage", bucketName), 0755)
}

// UploadFile enregistre un fichier dans le bucket spécifié.
func UploadFile(bucketName, fileName string, data []byte) error {
	// Construire le chemin complet du fichier
	filePath := filepath.Join("storage", bucketName, fileName)
	// Écrire les données dans le fichier
	return os.WriteFile(filePath, data, 0644)
}

// ListFiles liste tous les fichiers dans un bucket.
func ListFiles(bucketName string) ([]os.DirEntry, error) {
	return os.ReadDir(filepath.Join("storage", bucketName))
}

// DownloadFile récupère le contenu d'un fichier depuis un bucket.
func DownloadFile(bucketName, fileName string) ([]byte, error) {
	filePath := filepath.Join("storage", bucketName, fileName)
	return os.ReadFile(filePath)
}

// DeleteFile supprime un fichier d'un bucket.
func DeleteFile(bucketName, fileName string) error {
	filePath := filepath.Join("storage", bucketName, fileName)
	return os.Remove(filePath)
}



