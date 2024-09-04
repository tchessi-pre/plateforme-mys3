package handlers

import (
	"io/ioutil"
	"mys3/storage"
	"net/http"
	"strings"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	bucketName := r.URL.Query().Get("bucket")
	fileName := r.URL.Query().Get("file")
	
	if bucketName == "" || fileName == "" {
		http.Error(w, "Bucket and file name are required", http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}

	err = storage.UploadFile(bucketName, fileName, data)
	if err != nil {
		http.Error(w, "File upload failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ListFiles(w http.ResponseWriter, r *http.Request) {
	bucketName := r.URL.Query().Get("bucket")
	if bucketName == "" {
		http.Error(w, "Bucket name is required", http.StatusBadRequest)
		return
	}

	files, err := storage.ListFiles(bucketName)
	if err != nil {
		http.Error(w, "Could not list files", http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		w.Write([]byte(file.Name() + "\n"))
	}
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	bucketName := r.URL.Query().Get("bucket")
	fileName := r.URL.Query().Get("file")
	
	if bucketName == "" || fileName == "" {
		http.Error(w, "Bucket and file name are required", http.StatusBadRequest)
		return
	}

	data, err := storage.DownloadFile(bucketName, fileName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	bucketName := r.URL.Query().Get("bucket")
	fileName := r.URL.Query().Get("file")
	
	if bucketName == "" || fileName == "" {
		http.Error(w, "Bucket and file name are required", http.StatusBadRequest)
		return
	}

	err := storage.DeleteFile(bucketName, fileName)
	if err != nil {
		http.Error(w, "File deletion failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DownloadObject(w http.ResponseWriter, r *http.Request) {
	// Extraire le nom du bucket depuis l'URL
	bucketName := strings.TrimPrefix(r.URL.Path, "/buckets/")
	// Extraire le nom de l'objet depuis les paramètres de la requête
	objectName := r.URL.Query().Get("object")

	if bucketName == "" || objectName == "" {
		http.Error(w, "Bucket and object name are required", http.StatusBadRequest)
		return
	}

	// Télécharger l'objet en utilisant la fonction du fichier storage.go
	data, err := storage.DownloadFile(bucketName, objectName)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Configurer les en-têtes de la réponse pour le téléchargement
	w.Header().Set("Content-Disposition", "attachment; filename="+objectName)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}


