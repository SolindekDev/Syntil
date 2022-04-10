package main

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func generateToken() string {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	constants := "abcdefghijklmnoprstuwxqkjsadjJDNSJKADNASJDIKOQNDASJDM01934820138912R50-1MXKASDASJDKSANM-=-=;''LP['M,1OJMI"
	r := ""
	r += timestamp[0:3]
	for i := 0; i < 48; i++ {
		r += string(constants[rand.Intn(len(constants)-0) + 0])
	}
	r += timestamp[len(timestamp)-3:len(timestamp)]
	return r
}

func generateID() string {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	constants := "012345678992648100"
	r := ""
	r += timestamp[len(timestamp)-3:len(timestamp)]
	for i := 0; i < 10; i++ {
		r += string(constants[rand.Intn(len(constants)-0) + 0])
	}
	return r
}