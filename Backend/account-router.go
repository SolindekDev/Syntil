package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func loginPost(c *gin.Context) {
	email, emailExists := c.GetPostForm("email")
	password, passwordExists := c.GetPostForm("password")

	if emailExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Email not given" }); return }
	if passwordExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Password not given" }); return }

	filter := bson.D{{"email",email},{"password",password}}
	findLoginUser := FindUser(database, filter)

	if findLoginUser.Email == email && findLoginUser.Password == password {
		c.JSON(http.StatusOK, gin.H{ "status":"200", "message":"Found account", "token":findLoginUser.Token, "id":findLoginUser.ID })
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Email or Password is incorrect" })
		return
	}
}

func followPost(c *gin.Context) {
	token := c.Param("TOKEN")
	profile_id:= c.Param("PROFILE_ID")

	filter := bson.D{{"token",token}}
	findLoginUser := FindUser(database, filter)

	if findLoginUser.Token == token {

	}
}

func editPost(c *gin.Context) {
	email, emailExists := c.GetPostForm("email")
	token, tokenExists := c.GetPostForm("token")
	password, passwordExists := c.GetPostForm("password")
	newValue, newValueExists := c.GetPostForm("new_value")
	mode := c.Param("MODE")

	if emailExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Email not given" }); return }
	if passwordExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Password not given" }); return }
	if tokenExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Token not given" }); return }
	if newValueExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"New value of the update not given" }); return }

	filter := bson.D{{"email",email},{"password",password},{"token",token}}
	findLoginUser := FindUser(database, filter)

	if findLoginUser.Email == email && findLoginUser.Password == password && findLoginUser.Token == token {
		if mode == "email" {
			database.Collection("users").FindOneAndUpdate(context.TODO(), filter, bson.M{"$set":bson.M{"email":newValue}})
			c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Account has been updated successfully"})
			return
		} else if mode == "password" {
			database.Collection("users").FindOneAndUpdate(context.TODO(), filter, bson.M{"$set":bson.M{"password":newValue}})
			c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Account has been updated successfully"})
			return
		} else if mode == "username" {
			database.Collection("users").FindOneAndUpdate(context.TODO(), filter, bson.M{"$set":bson.M{"username":newValue}})
			c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Account has been updated successfully"})
			return
		} else if mode == "biographie" {
			database.Collection("users").FindOneAndUpdate(context.TODO(), filter, bson.M{"$set":bson.M{"biographie":newValue}})
			c.JSON(http.StatusOK, gin.H{"status": "200", "message": "Account has been updated successfully"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Accout has not been founded" })
		return
	}
}

func deletePost(c *gin.Context) {
	email, emailExists := c.GetPostForm("email")
	token, tokenExists := c.GetPostForm("token")
	password, passwordExists := c.GetPostForm("password")

	if emailExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Email not given" }); return }
	if passwordExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Password not given" }); return }
	if tokenExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Token not given" }); return }

	filter := bson.D{{"email",email},{"password",password},{"token",token}}
	findLoginUser := FindUser(database, filter)

	if findLoginUser.Email == email && findLoginUser.Password == password && findLoginUser.Token == token {
		database.Collection("users").FindOneAndDelete(context.TODO(), filter)
		c.JSON(http.StatusOK, gin.H{ "status":"200", "message":"Account has been deleted" })
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{ "status":"400", "message":"Accout has not been founded and as well deleted" })
		return
	}
}

func registerPost(c *gin.Context) {
	email, emailExists := c.GetPostForm("email")
	username, usernameExists := c.GetPostForm("username")
	password, passwordExists := c.GetPostForm("password")

	if emailExists != true { c.JSON(http.StatusBadRequest, gin.H{ "status": "400", "message": "Email not given!" }); return
	} else if usernameExists != true { c.JSON(http.StatusBadRequest, gin.H{
		"status":  "400",
		"message": "Username not given!",
	}); return
	} else if passwordExists != true { c.JSON(http.StatusBadRequest, gin.H{"status": "400", "message": "Password not given!" }); return }

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