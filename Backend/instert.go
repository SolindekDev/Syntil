package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func InsertUser(Database *mongo.Database, Email string, Username string, Password string, Token string, ID string, Biographie string, Followers []UserModel, AvatarURL string,  Timestamp time.Time) {
	user := UserModel{Email, Username, Password, Token, ID, Biographie, Followers, AvatarURL, Timestamp}
	collection := Database.Collection("users")
	_, e := collection.InsertOne(context.TODO(), user)
	if e != nil {
		fmt.Printf("Unknown error while inserting an user into the database")
	}
}

func InsertPost(Database *mongo.Database, Content string, ID string, AuthorID string, Likes []UserModel, Timestamp time.Time) {
	post := PostModel{Content, ID, AuthorID, Likes, Timestamp}
	collection := Database.Collection("posts")
	_, e := collection.InsertOne(context.TODO(), post)
	if e != nil {
		fmt.Printf("Unknown error while inserting an post into the database")
	}
}