package services

import (
	"context"
	"os"
	"time"
	m "valorant-agents/mongo"
	"valorant-agents/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var databaseName string = "valgents"
var playerCollection string = "players"
func CreatePlayer(player RequestPlayer) (Player, error) {
	client, err := m.ConnectToMongoDB()
	defer m.DisconnectFromMongoDB()
	if err != nil {
		return Player{}, err
	}
	coll := client.Database(databaseName).Collection(playerCollection)

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
	defer m.DisconnectFromMongoDB()
	if err != nil {
		return make([]Player, 0), err
	}
	coll := client.Database(databaseName).Collection(playerCollection)
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

func GetPlayerByID(ID string) (Player, error) {
	var player Player
	client, err := m.ConnectToMongoDB()
	defer m.DisconnectFromMongoDB()
	if err != nil {
		return player, nil
	}

	coll := client.Database(databaseName).Collection(playerCollection)
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return player, err
	}
	filter  := bson.D{{ "_id", objectID }}
	err = coll.FindOne(context.Background(), filter).Decode(&player)
	if err != nil {
		return player, nil
	}
	return player, nil
}

func CreatePlayerToken(player IDPlayerTokenClaim) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	tokenSecretKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &PlayerTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		Player: player,
	})
	return token.SignedString(tokenSecretKey)
}



func ReadPlayerToken(tokenString string) (IDPlayerTokenClaim, error) {
	playerClaim := IDPlayerTokenClaim{}
	err := godotenv.Load(".env")
	if err != nil {
		return playerClaim, err
	}
	tokenSecretKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &PlayerTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSecretKey, nil
	})

	if claims, ok := token.Claims.(*PlayerTokenClaims); ok && token.Valid {
		playerClaim = claims.Player
		return claims.Player, nil
	}  else {
		return playerClaim, err
	}
}