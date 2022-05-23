package db

import (
	"context"
	"github.com/ahloul/loyalty-reports/infrastructure/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewClient(db *Mongo) *mongo.Client {
	viper := config.NewViper()
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.Database.Connection))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	return client
}

type Mongo struct {
}

//ProvideInventoryRepo ... provide for the DI
func MongoC() Mongo {
	return Mongo{}
}
func (m *Mongo) ConnectPurchases() *mongo.Collection {
	viper := config.NewViper()

	clientOptions := options.Client().ApplyURI(viper.Database.Connection)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("reports").Collection("purchases")
	return collection
}
