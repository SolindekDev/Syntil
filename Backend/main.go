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

	r.POST("/api/1v/account/register", accountRegisterPost)
	r.POST("/api/1v/account/login", accountLoginPost)
	r.POST("/api/1v/account/delete", accountDeletePost)
	r.POST("/api/1v/account/update/:MODE", accountEditPost)
	r.GET("/api/1v/account/follow/:TOKEN/:PROFILE_ID", accountFollowPost)
	r.GET("/api/1v/account/find/id/:USER_ID", accountGetUserInfoPostByID)
	r.GET("/api/1v/account/find/token/:USER_TOKEN", accountGetUserInfoPostByToken)

	r.POST("/api/1v/post/create", postCreatePost)
	r.POST("/api/1v/post/delete", postDeletePost)
	r.POST("/api/1v/post/edit", postEditPost)
	r.GET("/api/1v/post/find/id/:POST_ID", postGetPostInfoGetByID)
	r.GET("/api/1v/post/like/:TOKEN/:POST_ID", postLikeGet)
	r.GET("/api/1v/post/allposts/:USER_ID/", allPostsOfUserGet)

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
