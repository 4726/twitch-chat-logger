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
	client                       *mongo.Client
	addr, dbName, collectionName string
}

func New(addr, dbName, collectionName string) *Storage {
	return &Storage{
		addr:           addr,
		dbName:         dbName,
		collectionName: collectionName,
	}
}

func (s *Storage) Connect() error {
	opts := options.Client().ApplyURI(s.addr)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}
	s.client = client

	return s.client.Ping(context.Background(), nil)
}

func (s *Storage) Add(cm storage.ChatMessage) error {
	collection := s.client.Database(s.dbName).Collection(s.collectionName)
	_, err := collection.InsertOne(context.Background(), cm)
	return err
}

func (s *Storage) Query(opts storage.QueryOptions) ([]storage.ChatMessage, error) {
	var res []storage.ChatMessage
	findOptions := options.Find()
	findOptions.SetLimit(1000)
	filter := createFilter(opts)
	collection := s.client.Database(s.dbName).Collection(s.collectionName)
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return res, err
	}
	if err := cursor.All(context.Background(), &res); err != nil {
		return res, err
	}
	return res, nil
}

func createFilter(opts storage.QueryOptions) bson.M {
	filter := bson.M{}
	if opts.SubscribeMin > 0 {
		filter["subscriber"] = true
		filter["subscribemonths"] = bson.M{"$gte": opts.SubscribeMin}
	}
	if opts.Term != "" {
		filter["message"] = fmt.Sprintf("/%v/", opts.Term)
	}
	if opts.Name != "" {
		filter["name"] = opts.Name
	}
	if opts.Channel != "" {
		filter["channel"] = opts.Channel
	}
	if !opts.Date.IsZero() {
		date := opts.Date
		startDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
		endDate := time.Date(date.Year(), date.Month(), date.Day()+1, 0, 0, 0, 0, date.Location())
		filter["time"] = bson.M{
			"$gte": startDate.Unix(),
			"$lte": endDate.Unix(),
		}
	}
	if opts.Admin {
		filter["admin"] = true
	}
	if opts.GlobalMod {
		filter["globalmod"] = true
	}
	if opts.Staff {
		filter["staff"] = true
	}
	if opts.Turbo {
		filter["turbo"] = true
	}
	if opts.BitsMin != 0 && opts.BitsMax != 0 {
		filter["bits"] = bson.M{
			"$gte": opts.BitsMin,
			"$lte": opts.BitsMax,
		}
	}

	return filter
}

func (s *Storage) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	return s.client.Disconnect(ctx)
}
