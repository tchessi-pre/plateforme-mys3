package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DeleteBucket est un gestionnaire pour supprimer un bucket sans authentification
func DeleteBucket(w http.ResponseWriter, r *http.Request) {
	// Log de la requête
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)

	// Accepter uniquement la méthode DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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

	// Vérifier si le bucket existe
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		http.Error(w, fmt.Sprintf("Bucket %s does not exist", bucketName), http.StatusNotFound)
		return
	}

	// Supprimer le bucket
	err = os.RemoveAll(bucketPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete bucket: %v", err), http.StatusInternalServerError)
		return
	}

	// Log de suppression du bucket
	fmt.Printf("Bucket %s deleted successfully\n", bucketName)

	// Répondre avec un code 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
