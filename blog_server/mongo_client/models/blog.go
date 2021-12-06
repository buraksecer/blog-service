package models

type Blog struct {
	Id       int    `bson:"_id,omitempty"`
	AuthorId string `bson:"author_id"`
	Content  string `bson:"content"`
	Title    string `bson:"title"`
}
