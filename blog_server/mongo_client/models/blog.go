package models

type Blog struct {
	Id       string `bson:"_id,omitempty"`
	AuthorId string `bson:"author_id"`
	Content  string `bson:"content"`
	Title    string `bson:"title"`
}
