package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorId string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}
