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

	r.POST("/api/1v/account/register", registerPost)
	r.POST("/api/1v/account/login", loginPost)
	r.POST("/api/1v/account/delete", deletePost)
	r.POST("/api/1v/account/update/:MODE", editPost)
	r.GET("/api/1v/account/follow/:TOKEN/:PROFILE_ID", followPost)
	r.GET("/api/1v/account/find/id/:USER_ID", getUserInfoPostByID)
	r.GET("/api/1v/account/find/token/:USER_TOKEN", getUserInfoPostByToken)

	r.POST("/api/1v/post/create", createPostAPI)

	r.GET("/", mainGet)

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
