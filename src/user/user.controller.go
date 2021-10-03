package user

import (
	"argoCD-golang/src/database/repositories"
	"argoCD-golang/src/database/schemas"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertUser(ctx *gin.Context) {
	users_collection := repositories.UsersCollection()

	var user_schema schemas.User

	err := ctx.ShouldBindJSON(&user_schema)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	insertedUser, err := users_collection.InsertOne(context.TODO(), &user_schema)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot create user: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"insertedUser": insertedUser,
		"user":         user_schema,
	})
}

func FindUsers(ctx *gin.Context) {
	users_collection := repositories.UsersCollection()
	findOptions := options.Find()

	var user_schema []*schemas.User
	users, err := users_collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot list Users: " + err.Error(),
		})

		return
	}

	for users.Next(context.TODO()) {
		var elem schemas.User
		err := users.Decode(&elem)
		if err != nil {
			ctx.JSON(500, gin.H{
				"Error": "Cannot Decode User: " + err.Error(),
			})

			return
		}

		user_schema = append(user_schema, &elem)
	}

	if err := users.Err(); err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot list Users: " + err.Error(),
		})

		return
	}

	// Close the cursor once finished
	users.Close(context.TODO())
	ctx.JSON(200, user_schema)
}

func UpdateUser(ctx *gin.Context) {
	users_collection := repositories.UsersCollection()

	userId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot parse user id: " + err.Error(),
		})

		return
	}

	var user_to_update interface{}

	err = ctx.ShouldBindJSON(&user_to_update)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	updateOptions := options.FindOneAndUpdate()
	// Retornar o objeto após ser atualizado
	updateOptions.SetReturnDocument(options.After)

	filter := bson.M{"_id": userId}
	update := bson.M{
		"$set": user_to_update,
	}

	var updatedUser schemas.User
	err = users_collection.FindOneAndUpdate(context.TODO(), filter, update, updateOptions).Decode(&updatedUser)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot update User: " + err.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"updated":   updatedUser,
		"to_update": user_to_update,
	})
}

func DeleteUser(ctx *gin.Context) {
	users_collection := repositories.UsersCollection()

	userId, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot parse user id: " + err.Error(),
		})

		return
	}

	filter := bson.M{"_id": userId}

	userDeleted, err := users_collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Error": "Cannot delete user: " + err.Error(),
		})

		return
	}

	if userDeleted.DeletedCount == 0 {
		ctx.JSON(404, "User not found")
	}

	ctx.JSON(200, "User deleted")
}
