package config

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db      *mongo.Database
	options *options.ClientOptions
}

var (
	onceMongo sync.Once
	db        MongoDB
)

func (m MongoDB) GetURI() string {
	return m.options.GetURI()
}

func (m MongoDB) GetDB() *mongo.Database {
	return m.db
}

func InitDatabase(addr, database string) MongoDB {
	fmt.Println("Connecting to MongoDB ...")
	onceMongo.Do(func() {
		db.options = options.Client().ApplyURI(addr)
		client, err := mongo.Connect(context.Background(), db.options)
		if err != nil {
			panic(err.Error())
		}

		db.db = client.Database(database)
		fmt.Printf("MongoDB Connected!\n\n")
	})

	return db
}
