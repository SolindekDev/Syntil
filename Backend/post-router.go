package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

func postDeletePost(c *gin.Context) {
	token, tokenExists := c.GetPostForm("token")
	postID, postIDExists := c.GetPostForm("post_id")

	if tokenExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Token not given" }); return }
	if postIDExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Post_ID not given" }); return }

	filter := bson.D{{"token",token}}
	findUser := FindUser(database, filter)

	if findUser.Token == token {
		filter = bson.D{{"id", postID}}
		post := FindPost(database, filter)
		if post.ID == postID {
			database.Collection("posts").FindOneAndDelete(context.TODO(), filter)
			c.JSON(http.StatusOK, gin.H{ "status":"200", "message":"Post has been deleted" })
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Post with this id has not been founded" })
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Account with this token has not been founded" })
		return
	}
}

func postLikeGet(c *gin.Context) {
	token := c.Param("TOKEN")
	postID := c.Param("POST_ID")

	filter := bson.D{{"token", token}}
	findUserByToken := FindUser(database, filter)

	if findUserByToken.Token != token {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Account with that token has not been founded" })
		return
	}

	filter = bson.D{{"id",postID}}
	findPost := FindPost(database, filter)

	if findPost.ID != postID {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Post with that id has not been founded" })
		return
	}

	alreadyLike := false
	for i := 0; i < len(findPost.Likes); i++ {
		if findPost.Likes[i].ID == findUserByToken.ID {
			alreadyLike = true
		}
	}

	if alreadyLike == true {
		database.Collection("posts").FindOneAndUpdate(context.TODO(), filter, bson.M{"$pull": bson.M{"likes": bson.M{"id": findUserByToken.ID}}})
		c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Successfully unliked the post"})
		return
	} else {
		database.Collection("posts").FindOneAndUpdate(context.TODO(), filter, bson.M{
			"$push": bson.M{
				"likes": bson.M {
					"username": findUserByToken.Username,
					"id": findUserByToken.ID,
					"biographie": findUserByToken.Biographie,
					"avatarurl": findUserByToken.AvatarURL,
				},
		}})
		c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Successfully liked the post"})
		return
	}
}

func allPostsOfUserGet(c *gin.Context) {
	userID := c.Param("USER_ID")

	filter := bson.D{{"id",userID}}
	findUser := FindUser(database, filter)

	if findUser.ID != userID {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"User with that id has not been founded" })
		return
	}

	filter = bson.D{{"authorid",userID}}
	cursor, err := database.Collection("posts").Find(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Unknown error while searching for post of user with id: " + userID })
		return
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	//for _, result := range results {
	//	fmt.Println(result)
	//}
	c.JSON(http.StatusOK, gin.H{ "status":"200", "message":"Found all post of this user", "posts":results})
}

func postGetPostInfoGetByID(c *gin.Context) {
	postID := c.Param("POST_ID")

	filter := bson.D{{"id",postID}}
	findPost := FindPost(database, filter)

	if findPost.ID == postID {
		filter = bson.D{{"id", findPost.AuthorID}}
		userAuthor := FindUser(database,filter)
		fmt.Println(userAuthor)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":"400",
			"message":"Post has found",
			"post": gin.H{
				"content": findPost.Content,
				"id": findPost.ID,
				"author":gin.H{
					"username": userAuthor.Username,
					"id": userAuthor.ID,
					"biographie": userAuthor.Biographie,
					"avatarurl": userAuthor.AvatarURL,
					"timestamp": userAuthor.Timestamp,
					"followers": userAuthor.Followers,
				},
				"likes": findPost.Likes,
				"timestamp": findPost.Timestamp,
			},
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Failed while searching for post with this id." })
		return
	}
}

func postEditPost(c *gin.Context) {
	content, contentExists := c.GetPostForm("content")
	postID, postIDExists := c.GetPostForm("post_id")
	token, tokenExists := c.GetPostForm("token")

	if contentExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"New content of the post is not given." }); return }
	if tokenExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Token is not given." }); return }
	if postIDExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"ID of the post is not given." }); return }

	if len(content) > 150 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "New content is too long! Maximum 150 characters." }); return }
	if len(content) < 1 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "New content is too short! Minimum 1 character." }); return }

	filter := bson.D{{"token",token}}
	findUser := FindUser(database, filter)

	if findUser.Token == token {
		filter = bson.D{{"id", postID}}
		post := FindPost(database, filter)
		if post.ID == postID {
			if post.AuthorID == findUser.ID {
				updateFilter := bson.M{"$set":bson.M{"content":content}}
				database.Collection("posts").FindOneAndUpdate(context.TODO(), filter, updateFilter)
				c.JSON(http.StatusOK, gin.H{ "status":"200", "message":"Post has been updated" })
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Your are not the author of the post" })
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Post with this id has not been founded" })
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Account with this token has not been founded" })
		return
	}
}

func postCreatePost(c *gin.Context) {
	content, contentExists := c.GetPostForm("content")
	token, tokenExists := c.GetPostForm("token")

	if contentExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Content of the post is not given." }); return }
	if tokenExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Token is not given." }); return }

	if len(content) > 150 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Content is too long! Maximum 150 characters." }); return }
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