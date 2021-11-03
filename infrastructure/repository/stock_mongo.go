package repository

import (
	"context"
	"time"

	"github.com/riquellopes/closed-stock-consumer/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoDataBase() *mongo.Database {
	ctx, cancel := mongoContext()
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(""))

	if err != nil {
		panic(err)
	}

	if err = client.Connect(ctx); err != nil {
		panic(err)
	}

	database := client.Database("")
	return database
}

func mongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

type StockMongo struct {
	Collection *mongo.Collection
}

func (repository *StockMongo) Save(stock entity.Stock) {
	ctx, cancel := mongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{})

	if err != nil {
		panic(err)
	}
}
