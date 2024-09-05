package handlers

import (
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// loadMinioClient initialise un client MinIO à partir des variables d'environnement
func loadMinioClient() (*minio.Client, error) {
	// Récupération des variables d'environnement
	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")

	fmt.Println("MINIO_ENDPOINT:", minioEndpoint)
fmt.Println("MINIO_ACCESS_KEY:", accessKey)
fmt.Println("MINIO_SECRET_KEY:", secretKey)


	// Vérification que toutes les variables d'environnement sont définies
	if minioEndpoint == "" || accessKey == "" || secretKey == "" {
		return nil, fmt.Errorf("missing required environment variables: MINIO_ENDPOINT, MINIO_ACCESS_KEY, MINIO_SECRET_KEY")
	}

	// Initialiser le client MinIO
	minioClient, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // Mettre à true si vous utilisez HTTPS
	})
	if err != nil {
		return nil, fmt.Errorf("unable to initialize MinIO client: %w", err)
	}

	return minioClient, nil
}

