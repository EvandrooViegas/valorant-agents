package services

import (
	"context"
	"valorant-agents/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


const (
	COMMENT_COLLECTION_NAME = "comments"
)

func GetComments() ([]Comment, error) {
	var comments []Comment 
	client, err := mongo.ConnectToMongoDB()
	if err != nil {
		return make([]Comment, 0), nil
	}
	coll := client.Database(DATABASE_NAME).Collection(COMMENT_COLLECTION_NAME)
	filter := bson.D{{ "parent_id", "" }}
	cursor, err := coll.Find(context.Background(), filter)
	err = cursor.All(context.TODO(), &comments) 
	if err != nil {
		return make([]Comment, 0), err
	}
	return comments, nil
}

func GetCommentReplies(parentID string) ([]Comment, error) {
	var comments []Comment 
	client, err := mongo.ConnectToMongoDB()
	if err != nil {
		return make([]Comment, 0), nil
	}
	coll := client.Database(DATABASE_NAME).Collection(COMMENT_COLLECTION_NAME)
	filter := bson.D{{ "parent_id", parentID }}
	cursor, err := coll.Find(context.Background(), filter)
	err = cursor.All(context.TODO(), &comments) 
	if err != nil {
		return make([]Comment, 0), err
	}
	return comments, nil
}

func CreateComment(comment DatabaseComment) (Comment, error) {
	client, err := mongo.ConnectToMongoDB()
	defer mongo.DisconnectFromMongoDB()
	if err != nil {
		return Comment{}, err
	}

	coll := client.Database(DATABASE_NAME).Collection(COMMENT_COLLECTION_NAME)
	result, err := coll.InsertOne(context.TODO(), comment)
	if err != nil {
		return Comment{}, err
	}
	ID := result.InsertedID.(primitive.ObjectID).Hex()
	return Comment{
		ID: ID,
		AgentID: comment.AgentID,
		ParentID: comment.ParentID,
		Text: comment.Text,
		Replies: make([]Comment, 0),
		Author: comment.Author,
	}, nil

}