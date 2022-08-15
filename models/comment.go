package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Content  string             `json:"content,omitempty"`
	Upvote   int                `json:"upvote,omitempty"`
	Downvote int                `json:"downvote,omitempty"`
}
