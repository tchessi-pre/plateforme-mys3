package handlers

import (
	"mys3/storage"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Bucket struct {
	Name string `xml:"Name"`
}

func CreateBucket(w http.ResponseWriter, r *http.Request) {
	// Lire le corps de la requête
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parser le XML
	var bucket Bucket
	err = xml.Unmarshal(body, &bucket)
	if err != nil {
		http.Error(w, "Invalid XML", http.StatusBadRequest)
		return
	}

	// Vérifier que le nom du bucket est présent
	if bucket.Name == "" {
		http.Error(w, "Bucket name is required", http.StatusBadRequest)
		return
	}

	// Créer le bucket
	err = storage.CreateBucket(bucket.Name)
	if err != nil {
		http.Error(w, "Bucket creation failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UploadObject(w http.ResponseWriter, r *http.Request) {
	// Extraire le nom du bucket depuis l'URL
	bucketName := strings.TrimPrefix(r.URL.Path, "/buckets/")
	// Extraire le nom de l'objet depuis les paramètres de la requête
	objectName := r.URL.Query().Get("object")
	
	if bucketName == "" || objectName == "" {
		http.Error(w, "Bucket and object name are required", http.StatusBadRequest)
		return
	}

	// Lire le contenu du corps de la requête (le contenu de l'objet)
	objectData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// Uploader l'objet en utilisant la fonction du fichier storage.go
	err = storage.UploadFile(bucketName, objectName, objectData)
	if err != nil {
		http.Error(w, "Failed to upload object", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ListBuckets(w http.ResponseWriter, r *http.Request) {
	// Ouvrir le répertoire de stockage
	bucketsDir := "storage"
	files, err := os.ReadDir(bucketsDir)
	if err != nil {
		http.Error(w, "Could not list buckets", http.StatusInternalServerError)
		return
	}

	// Lister les répertoires (buckets)
	for _, file := range files {
		if file.IsDir() {
			w.Write([]byte(file.Name() + "\n"))
		}
	}
}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	// Extraire le nom du bucket depuis l'URL
	bucketName := strings.TrimPrefix(r.URL.Path, "/buckets/")
	// Extraire le nom de l'objet depuis les paramètres de la requête
	objectName := r.URL.Query().Get("object")

	if bucketName == "" || objectName == "" {
		http.Error(w, "Bucket and object name are required", http.StatusBadRequest)
		return
	}

	// Supprimer l'objet en utilisant la fonction du fichier storage.go
	err := storage.DeleteFile(bucketName, objectName)
	if err != nil {
		http.Error(w, "File deletion failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Object deleted successfully"))
}

func DeleteBucket(w http.ResponseWriter, r *http.Request) {
	// Extraire le nom du bucket depuis l'URL
	bucketName := strings.TrimPrefix(r.URL.Path, "/buckets/")

	if bucketName == "" {
		http.Error(w, "Bucket name is required", http.StatusBadRequest)
		return
	}

	// Construire le chemin du bucket
	bucketPath := filepath.Join("storage", bucketName)

	// Vérifier si le bucket existe
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		http.Error(w, "Bucket not found", http.StatusNotFound)
		return
	}

	// Supprimer le bucket (et tous les fichiers qu'il contient)
	err := os.RemoveAll(bucketPath)
	if err != nil {
		http.Error(w, "Failed to delete bucket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bucket deleted successfully"))
}