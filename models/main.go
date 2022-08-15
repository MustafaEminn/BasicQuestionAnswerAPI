package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Main struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Title    string             `json:"title,omitempty"`
	Content  string             `json:"content,omitempty"`
	Upvote   int                `json:"upvote,omitempty"`
	Downvote int                `json:"downvote,omitempty"`
	Comments []Comment          `json:"comments,omitempty"`
}
