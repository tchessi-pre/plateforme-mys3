# Utilise une image de base Go officielle
FROM golang:1.19-alpine

# Créer un répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers Go dans le répertoire de travail
COPY . .

# Construire l'application
RUN go build -o mycli .

# Exposer le port que l'application utilise (facultatif)
EXPOSE 8080

# Commande par défaut
CMD ["./mycli"]
    