package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// checkCredentials vérifie les informations d'identification envoyées avec la requête
func checkCredentials(r *http.Request) bool {
	// Autoriser les requêtes sur les chemins "/probe-bsign" sans authentification
	if strings.HasPrefix(r.URL.Path, "/probe-bsign") || r.URL.Path == "/" {
		return true
	}

	// Récupérer les informations d'identification envoyées par le client (Basic Auth)
	user, pass, ok := r.BasicAuth()
	if !ok {
		fmt.Println("No credentials provided")
		return false
	}

	// Log des informations d'identification pour déboguer
	fmt.Printf("Received credentials: user=%s, pass=%s\n", user, pass)

	// Comparer avec les informations d'identification attendues
	if user == "admin" && pass == "admin123" {
		return true
	}
	fmt.Println("Invalid credentials")
	return false
}

// CreateBucket est un gestionnaire pour créer un bucket compatible avec l'API S3
func CreateBucket(w http.ResponseWriter, r *http.Request) {
	// Vérifier les informations d'identification, sauf pour les chemins spéciaux
	if !checkCredentials(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Log de la requête
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)

	// Vérifier si la requête est pour "/probe-bsign" ou une autre requête de santé
	if strings.HasPrefix(r.URL.Path, "/probe-bsign") || r.URL.Path == "/" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
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

		// Simuler la création du bucket
		fmt.Printf("Bucket %s created successfully\n", bucketName)

		// Répondre avec le header Location et un code 200 OK
		w.Header().Set("Location", fmt.Sprintf("/%s", bucketName))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Bucket %s created successfully", bucketName)))

	case http.MethodHead, http.MethodGet:
		// Ces méthodes peuvent être utilisées pour vérifier la connectivité
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))

	default:
		// Retourner une erreur 405 si la méthode n'est pas autorisée
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
