package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUser(database *mongo.Database, filter bson.D) UserModel {
	var res UserModel
	database.Collection("users").FindOne(context.TODO(), filter).Decode(&res)
	return res
}

func FindPost(database *mongo.Database, filter bson.D) PostModel {
	var res PostModel
	database.Collection("posts").FindOne(context.TODO(), filter).Decode(&res)
	return res
}