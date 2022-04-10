package main

import "time"

type UserModel struct {
	Email string
	Username string
	Password string
	Token string
	ID string
	Biographie string
	Followers []UserModel
	AvatarURL string
	Timestamp time.Time
}

type PostModel struct {
	Content string
	ID string
	AuthorID string
	Likes []UserModel
	Timestamp time.Time
}