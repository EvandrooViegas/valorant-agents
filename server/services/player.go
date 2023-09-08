package services

import (
	"context"
	"fmt"
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

const (
	PLAYER_COLLECTION = "players"
)

var TOKEN_AND_COOKIE_EXPIRITY_TIME = time.Now().Add(720 * time.Hour) // 1 Month
func CreatePlayer(player RegisterPlayerRequest) (Player, error) {
	client, err := m.ConnectToMongoDB()
	defer m.DisconnectFromMongoDB()
	if err != nil {
		return Player{}, err
	}
	coll := client.Database(DATABASE_NAME).Collection(PLAYER_COLLECTION)
	hashedPassword := utils.CreateHash(player.Password)
	result, err := coll.InsertOne(context.TODO(), DatabaseNewPlayer{
		Username:    player.Username,
		Description: player.Description,
		Password:    hashedPassword,
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
		Password:    hashedPassword,
		Description: player.Description,
	}, nil
}

func FilterPlayerWithUsername(username string) ([]Player, error) {
	client, err := m.ConnectToMongoDB()
	defer m.DisconnectFromMongoDB()
	if err != nil {
		return make([]Player, 0), err
	}
	coll := client.Database(DATABASE_NAME).Collection(PLAYER_COLLECTION)
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

	coll := client.Database(DATABASE_NAME).Collection(PLAYER_COLLECTION)
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return player, err
	}
	filter := bson.D{{"_id", objectID}}
	err = coll.FindOne(context.Background(), filter).Decode(&player)
	if err != nil {
		return player, nil
	}
	return player, nil
}

func GetPlayerByUsername(username string) (Player, error) {
	var player Player
	client, err := m.ConnectToMongoDB()
	defer m.DisconnectFromMongoDB()
	if err != nil {
		return player, nil
	}

	coll := client.Database(DATABASE_NAME).Collection(PLAYER_COLLECTION)
	filter := bson.D{{"username", username}}
	err = coll.FindOne(context.Background(), filter).Decode(&player)
	if err != nil {
		return player, nil
	}
	return player, nil
}

func CreatePlayerToken(ID string) (string, error) {
	player := IDPlayerTokenClaim{ID}
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	tokenSecretKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &PlayerTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(TOKEN_AND_COOKIE_EXPIRITY_TIME),
		},
		Player: player,
	})
	return token.SignedString(tokenSecretKey)
}

func ReadTokenAndReturnPlayer(tokenString string) (Player, error) {
	claims, err := ReadPlayerToken(tokenString)
	if err != nil {
		return Player{}, err
	}

	ID := claims.ID
	player, err := GetPlayerByID(ID)
	if err != nil {
		return Player{}, err
	}
	return player, nil
}

func ReadPlayerToken(tokenString string) (IDPlayerTokenClaim, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return IDPlayerTokenClaim{}, err
	}
	tokenSecretKey := []byte(os.Getenv("TOKEN_SECRET_KEY"))

	token, err := jwt.ParseWithClaims(tokenString, &PlayerTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenSecretKey, nil
	})
	if err != nil {
		return IDPlayerTokenClaim{}, err
	}
	if claims, ok := token.Claims.(*PlayerTokenClaims); ok && token.Valid {
		return claims.Player, nil
	} else {
		return IDPlayerTokenClaim{}, err
	}
}

type RegisterPlayerReturn struct {
	Token         string
	ID            string
	AlreadyExists bool
}

func RegisterPlayer(player RegisterPlayerRequest) (RegisterPlayerReturn, error) {
	foundPlayers, err := FilterPlayerWithUsername(player.Username)
	if err != nil {
		return RegisterPlayerReturn{}, err
	}
	if len(foundPlayers) > 0 {
		return RegisterPlayerReturn{AlreadyExists: true}, nil
	}

	nPlayer, err := CreatePlayer(player)
	if err != nil {
		return RegisterPlayerReturn{}, err
	}
	token, err := CreatePlayerToken(nPlayer.ID)
	if err != nil {
		return RegisterPlayerReturn{}, err
	}
	return RegisterPlayerReturn{
		Token: token,
		ID:    nPlayer.ID,
	}, nil
}

func LoginPlayer(requestPlayer LoginPlayerRequest) (Player, error) {

	foundPlayer, err := GetPlayerByUsername(requestPlayer.Username)
	reqPlayerHashedPsw := utils.CreateHash(requestPlayer.Password)
	if err != nil {
		return Player{}, nil
	}
	// check if the found user and the login user credentials are the same
	if foundPlayer.Password == reqPlayerHashedPsw {
		return foundPlayer, nil
	} else {
		return Player{}, fmt.Errorf("Could not create the Player")
	}
}
