package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func registerPost(c *gin.Context) {
	email, emailExists := c.GetPostForm("email")
	username, usernameExists := c.GetPostForm("username")
	password, passwordExists := c.GetPostForm("password")

	if emailExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Email not given!", }); return
	} else if usernameExists != true { c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Username not given!",
		}); return
	} else if passwordExists != true { c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Password not given!",}); return }

	if len(email) < 6 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Email is too short! Minimum 6 characters." }); return }
	if len(email) > 80 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Email is too long! Maximum 80 characters." }); return }
	//if isEmailValid(email) == true { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Email is not correct" }); return }

	if len(username) < 3 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Username is too short! Minimum 3 characters." }); return }
	if len(username) > 32 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Username is too long! Maximum 32 characters." }); return }

	if len(password) < 8 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Password is too short! Minimum 8 characters." }); return }
	if len(password) > 60 { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Password is too long! Maximum 60 characters." }); return }

	filter := bson.D{{"email",email}}
	findThatSameUser := FindUser(database, filter)

	if findThatSameUser.Email == email { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Email is already used"}); return }

	filter = bson.D{{"username",username}}
	findThatSameUser = FindUser(database, filter)

	if findThatSameUser.Username == username { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Username is already used"}); return }

	var followers []UserModel
	id := generateID()
	token := generateToken()
	InsertUser(
		database,
		email,
		username,
		password,
		token,
		id,
		"no bio yet",
		followers,
		"/avatar/"+id+".png",
		time.Now(),
	)

	c.JSON(http.StatusOK, gin.H{"status":"200","message":"Account has been created","token":token})
}