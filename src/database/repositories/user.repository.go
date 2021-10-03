package repositories

import (
	"argoCD-golang/src/database"

	"go.mongodb.org/mongo-driver/mongo"
)

func UsersCollection() *mongo.Collection {
	users_collection := database.GetDatabase().Collection("users")

	return users_collection
}
