package handlers

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Structure pour le corps de la requête CreateBucket en XML
type CreateBucketConfiguration struct {
	XMLName            xml.Name `xml:"CreateBucketConfiguration"`
	LocationConstraint string   `xml:"LocationConstraint"`
}

// CreateBucket est un gestionnaire qui crée un bucket (compatible avec l'API S3)
func CreateBucket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupérer le nom du bucket depuis l'URL "Host: bucket.s3.amazonaws.com"
	host := r.Host
	bucketName := strings.Split(host, ".")[0] // Extraire le nom du bucket

	if bucketName == "" {
		http.Error(w, "Bucket name is required", http.StatusBadRequest)
		return
	}

	// Lire le corps de la requête (XML) si présent
	var bucketConfig CreateBucketConfiguration
	if r.ContentLength > 0 {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		// Parse du corps en XML
		if err := xml.Unmarshal(body, &bucketConfig); err != nil {
			http.Error(w, "Invalid XML in request body", http.StatusBadRequest)
			return
		}
	}

	// Logique pour créer un bucket (simulé ici)
	fmt.Printf("Bucket %s created with location constraint: %s\n", bucketName, bucketConfig.LocationConstraint)

	// Répondre avec le header Location et un code 200 OK
	w.Header().Set("Location", fmt.Sprintf("/%s", bucketName))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Bucket %s created successfully", bucketName)))
}
