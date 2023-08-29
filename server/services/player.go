package services

import (
	"context"
	m "valorant-agents/mongo"
	"valorant-agents/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePlayer(player RequestPlayer) (Player, error) {
	client, err := m.ConnectToMongoDB()
	defer m.DisconnectFromMongoDB()
	if err != nil {
		return Player{}, err
	}
	coll := client.Database("valgents").Collection("players")

	result, err := coll.InsertOne(context.TODO(), DatabaseNewPlayer{
		Username:    player.Username,
		Description: player.Description,
		Password:    utils.CreateHash(player.Password),
		Avatar:      player.Avatar,
	})
	if err != nil {
		return Player{}, err
	}
	nPlayerID := result.InsertedID.(primitive.ObjectID).Hex()

	return Player{
		ID:          nPlayerID,
		Avatar:      player.Avatar,
		Username:    player.Username,
		Description: player.Description,
	}, nil
}

func FilterPlayerWithUsername(username string) ([]Player, error) {
	client, err := m.ConnectToMongoDB()
	if err != nil {
		return make([]Player, 0), err
	}
	coll := client.Database("valgents").Collection("players")
	filter := bson.D{{Key: "username", Value: username}}

	var player Player
	err = coll.FindOne(context.TODO(), filter).Decode(&player)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return make([]Player, 0), nil
		}
		return make([]Player, 0), err
	}
	return []Player{player}, nil
}
