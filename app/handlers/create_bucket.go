package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// CreateBucket est un gestionnaire pour créer un bucket sans authentification
func CreateBucket(w http.ResponseWriter, r *http.Request) {
	// Log de la requête
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)

	// Vérifier si la requête est pour "/probe-bsign" ou une autre requête de santé
	if strings.HasPrefix(r.URL.Path, "/probe-bsign") || r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server S3 started successfully"))
		return
	}

	// Accepter plusieurs méthodes HTTP pour la gestion des buckets
	switch r.Method {
	case http.MethodPut:
		// Extraire le nom du bucket depuis l'URL
		bucketName := strings.TrimPrefix(r.URL.Path, "/")
		if bucketName == "" {
			http.Error(w, "Bucket name is required", http.StatusBadRequest)
			return
		}

		// Utiliser `os.Getwd()` pour obtenir le répertoire de travail courant
		workingDir, err := os.Getwd()
		if err != nil {
			http.Error(w, "Unable to determine working directory", http.StatusInternalServerError)
			return
		}

		// Définir le chemin absolu du dossier "storage" et du bucket
		storagePath := filepath.Join(workingDir, "storage") // Le dossier "storage" absolu
		bucketPath := filepath.Join(storagePath, bucketName)

		// Vérifier si le dossier "storage" existe, sinon le créer
		if _, err := os.Stat(storagePath); os.IsNotExist(err) {
			err = os.MkdirAll(storagePath, os.ModePerm) // Créer le dossier "storage"
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed to create storage directory: %v", err), http.StatusInternalServerError)
				return
			}
			fmt.Println("Storage directory created")
		}

		// Créer le dossier du bucket à l'intérieur du dossier "storage"
		err = os.MkdirAll(bucketPath, os.ModePerm)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create bucket directory: %v", err), http.StatusInternalServerError)
			return
		}

		// Log de création du bucket
		fmt.Printf("Bucket %s created successfully\n", bucketName)

		// Répondre avec le header Location et un code 200 OK
		w.Header().Set("Location", fmt.Sprintf("/%s", bucketName))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Bucket %s created successfully and directory created in storage", bucketName)))

	case http.MethodHead, http.MethodGet:
		// Ces méthodes peuvent être utilisées pour vérifier la connectivité
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	default:
		// Retourner une erreur 405 si la méthode n'est pas autorisée
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
