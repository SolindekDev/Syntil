package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var port = ":80"
var uri = "mongodb://localhost:27017/"
var database *mongo.Database

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/api/1v/register", registerPost)

	return r
}

func main() {
	database = connectToDatabase(uri)
	r := setupRouter()
	err := r.Run(port)
	if err != nil {
		fmt.Println("Unknown error while starting the HTTP server.")
	}
}

/*
var followers []UserModel
InsertUser(database, "Email@email.com", "Username123", "Password123", "SOKDOASODKSADNASKJN1IJWUI~!", "1234567890", "Random guy.", followers, "https://avatar.url/url.png", time.Now())
 */
