package collection

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func Getcollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("myGoappDB").Collection("Posts")
	return collection
}
