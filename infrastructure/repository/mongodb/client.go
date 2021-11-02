package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClosedDataBase struct {
	data mongo.Database
}

func DataBase() *ClosedDataBase {
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
	return &ClosedDataBase{data: *database}
}

func mongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
