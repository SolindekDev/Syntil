package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func deletePostAPI(c *gin.Context)

func createPostAPI(c *gin.Context) {
	content, contentExists := c.GetPostForm("content")
	token, tokenExists := c.GetPostForm("token")

	if contentExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Content of the post is not given." }); return }
	if tokenExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Token is not given." }); return }

	if len(content) < 150 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Content is too long! Maximum 150 characters." }); return }
	if len(content) < 1 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Content is too short! Minimum 1 character." }); return }

	filter := bson.D{{"token",token}}
	findUserAuthor := FindUser(database, filter)

	if findUserAuthor.Token == token {
		postID := generateID()
		likes := []UserModel{}
		accTimestamp := time.Now()
		InsertPost(database, content, postID, findUserAuthor.ID, likes, accTimestamp)

		c.JSON(http.StatusOK, gin.H{
			"status":"200",
			"message":"Post has been created.",
			"post": gin.H{
				"content": content,
				"id": postID,
				"author": gin.H{
					"username": findUserAuthor.Username,
					"id": findUserAuthor.ID,
					"biographie": findUserAuthor.Biographie,
					"avatarurl": findUserAuthor.AvatarURL,
				},
				"likes": likes,
				"timestamp": accTimestamp,
			},
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Accout with that token has not been founded" })
		return
	}
}