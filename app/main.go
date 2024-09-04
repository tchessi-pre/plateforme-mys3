package main

import (
	"log"
	"mys3/handlers"
	"net/http"
	"strings"
)

func main() {
	// Route pour créer un bucket (POST)
	http.HandleFunc("/create-bucket", handlers.CreateBucket)

	// Route pour uploader un fichier dans un bucket (PUT)
	http.HandleFunc("/upload-file", handlers.UploadFile)

	// Route pour lister les fichiers dans un bucket (GET)
	http.HandleFunc("/list-files", handlers.ListFiles)

	// Route pour télécharger un fichier spécifique (GET)
	http.HandleFunc("/download-file", handlers.DownloadFile)

	// Route pour supprimer un fichier spécifique (DELETE)
	http.HandleFunc("/delete-file", handlers.DeleteFile)

	// Route pour lister les buckets (GET) et créer un bucket (POST)
	http.HandleFunc("/buckets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.ListBuckets(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateBucket(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route pour les opérations sur des objets spécifiques dans un bucket et supprimer un bucket
	http.HandleFunc("/buckets/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.DownloadObject(w, r)
		} else if r.Method == http.MethodPut {
			handlers.UploadObject(w, r)
		} else if r.Method == http.MethodDelete {
			// On vérifie si l'URL est pour supprimer un bucket ou un objet
			if strings.Contains(r.URL.Path, "/buckets/") && r.URL.Query().Get("object") == "" {
				handlers.DeleteBucket(w, r)
			} else {
				handlers.DeleteObject(w, r) 
			}
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
