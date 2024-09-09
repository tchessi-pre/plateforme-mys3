package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// PutObject gère le téléchargement d'un objet dans un bucket
func PutObject(w http.ResponseWriter, r *http.Request) {
	// Assurez-vous que la méthode de requête est PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Extraire le nom du bucket et la clé de l'objet depuis l'URL
	urlParts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	if len(urlParts) != 2 {
		http.Error(w, "Le nom du bucket et la clé de l'objet sont requis", http.StatusBadRequest)
		return
	}
	bucketName := urlParts[0]
	objectKey := urlParts[1] // Ce sera le nom du fichier

	// Utiliser `os.Getwd()` pour obtenir le répertoire de travail courant
		workingDir, err := os.Getwd()
		if err != nil {
			http.Error(w, "Unable to determine working directory", http.StatusInternalServerError)
			return
		}
	// Afficher explicitement le nom du fichier
	fmt.Printf("Nom du fichier à télécharger : %s\n", objectKey)

	// Définir le chemin du dossier "storage" à partir du répertoire courant (répertoire projet)
	storagePath := filepath.Join(workingDir, "storage") // Répertoire relatif "storage" dans la structure du projet
	bucketPath := filepath.Join(storagePath, bucketName) // Chemin vers le bucket dans le dossier "storage"
	objectPath := filepath.Join(bucketPath, objectKey)   // Chemin vers le fichier dans le bucket

	// Vérifiez que le bucket existe
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		http.Error(w, "Le bucket n'existe pas", http.StatusNotFound)
		return
	}

	// Créer les répertoires pour l'objet s'ils n'existent pas
	objectDir := filepath.Dir(objectPath)
	if err := os.MkdirAll(objectDir, os.ModePerm); err != nil {
		http.Error(w, "Impossible de créer le répertoire de l'objet", http.StatusInternalServerError)
		return
	}

	// Créer le fichier où l'objet sera stocké
	file, err := os.Create(objectPath)
	if err != nil {
		http.Error(w, "Impossible de créer le fichier de l'objet", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Lire le corps de la requête (le fichier) et l'écrire dans le fichier
	bytesWritten, err := io.Copy(file, r.Body)
	if err != nil {
		http.Error(w, "Échec de l'écriture des données de l'objet", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Vérification du nombre d'octets écrits
	if bytesWritten == 0 {
		http.Error(w, "Aucune donnée reçue", http.StatusBadRequest)
		return
	}

	// Log pour indiquer que l'objet a été uploadé avec succès
	fmt.Printf("Fichier '%s' téléchargé dans le bucket '%s' (%d octets)\n", objectKey, bucketName, bytesWritten)

	// Réponse 200 OK
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Fichier '%s' téléchargé avec succès", objectKey)))
}
