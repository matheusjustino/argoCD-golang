package main

import (
	"argoCD-golang/src/database"
	"argoCD-golang/src/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.StartDB()
	server := server.NewServer()
	server.Run()
}

//  kubectl port-forward svc/argocd-golang-service 8080:8080
