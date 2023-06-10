package main

import (
	"github.com/gin-gonic/gin"
	routes "github.com/myusername/go_crud/crud"
)

func main() {
	router := gin.Default()
	router.POST("/", routes.CreatePost)
	router.GET("getOne/:postId", routes.ReadOnePost)
	router.Run()
}
