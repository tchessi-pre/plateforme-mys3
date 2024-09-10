package main

import (
	"log"
	"net/http"
	"mys3/handlers"
)

func main() {
	// Routes pour la gestion des buckets
	http.HandleFunc("/", handlers.CreateBucket) // Créer un bucket
	http.HandleFunc("/list-files", handlers.ListFiles)       // Lister les fichiers d'un bucket

	// Routes pour la gestion des fichiers dans les buckets
	http.HandleFunc("/upload-file", handlers.UploadFile)     // Uploader un fichier
	http.HandleFunc("/download-file", handlers.DownloadFile) // Télécharger un fichier
	http.HandleFunc("/delete-file", handlers.DeleteFile)     // Supprimer un fichier

	// Gérer les objets dans un bucket avec PutObject (Upload d'objets dans un bucket via une URL)
	http.HandleFunc("/put-object/", handlers.PutObject)

	// Démarrer le serveur sur le port 8080
	log.Println("Server S3 started on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
