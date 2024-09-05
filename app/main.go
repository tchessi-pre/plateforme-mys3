package main

import (
	"log"
	"net/http"
	"mys3/handlers"
)

func main() {
	// Routes compatibles avec S3
	http.HandleFunc("/create-bucket", handlers.CreateBucket) // Créer un bucket
	http.HandleFunc("/upload-file", handlers.UploadFile)     // Uploader un fichier
	http.HandleFunc("/delete-file", handlers.DeleteFile)     // Supprimer un fichier
	http.HandleFunc("/list-files", handlers.ListFiles)       // Lister les fichiers
	http.HandleFunc("/download-file", handlers.DownloadFile) // Télécharger un fichier

	// Démarrer le serveur sur le port 8080
	log.Println("My API S3 en cours d'exécution sur le port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
