package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/4726/twitch-chat-logger/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	client *mongo.Client
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect() error {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}
	s.client = client

	return s.client.Ping(context.Background(), nil)
}

func (s *Storage) Add(cm storage.ChatMessage) error {
	collection := s.client.Database("db").Collection("messages")
	_, err := collection.InsertOne(context.Background(), cm)
	return err
}

func (s *Storage) Query(channel, term, name string, date time.Time) ([]storage.ChatMessage, error) {
	var res []storage.ChatMessage
	findOptions := options.Find()
	findOptions.SetLimit(1000)
	filter := bson.M{}
	if term != "" {
		filter["message"] = fmt.Sprintf("/%v/", term)
	}
	if name != "" {
		filter["name"] = name
	}
	if channel != "" {
		filter["channel"] = channel
	}
	if !date.IsZero() {
		startDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
		endDate := time.Date(date.Year(), date.Month(), date.Day()+1, 0, 0, 0, 0, date.Location())
		filter["time"] = bson.M{
			"$gte": startDate.Unix(),
			"$lte": endDate.Unix(),
		}
	}
	collection := s.client.Database("db").Collection("messages")
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return res, err
	}
	if err := cursor.All(context.Background(), &res); err != nil {
		return res, err
	}
	return res, nil
}
