package crud

import (
	"context"
	"log"
	"net/http"
	"time"

	collection "github.com/myusername/go_crud/collection"
	database "github.com/myusername/go_crud/databases"
	model "github.com/myusername/go_crud/model"

	"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *gin.Context) {
	var DB = database.ConnectDB()
	var postCollection = collection.Getcollection(DB, "Posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	post := new(model.Posts)
	defer cancel()

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		log.Fatal(err)
		return
	}

	postPayload := model.Posts{
		//Id:        post.ID,
		Title:   post.Title,
		Article: post.Article,
	}

	result, err := postCollection.InsertOne(ctx, postPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Posted successfully", "Data": map[string]interface{}{"data": result}})
}
