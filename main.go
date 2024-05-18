package main

import (
	"apigo/database"
	"apigo/posts"
	"apigo/users"

	"github.com/gin-gonic/gin"
)

func init() {
	database.InitDatabase("./test.db")
}

func main() {
	userHandler := users.MakeUserHandle()
	postModel := posts.MakePostModel()
	postModel.Migrate()

	router := gin.Default()
	router.POST("/users/signup", userHandler.SignUpHandler)
	router.POST("/users/login", userHandler.LoginHandler)
	router.GET("/users/", userHandler.GetUsersHandler)
	router.Run()
}
