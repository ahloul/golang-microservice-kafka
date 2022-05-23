package repository

import (
	"context"
	"github.com/ahloul/loyalty-reports/infrastructure/db"
	"github.com/ahloul/loyalty-reports/models"
	_ "github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type PurchaseRepository struct {
	clinet *db.Mongo
}

//func NewPurchaseRepository(db *gorm.DB, rd *redis.Client) subject.Repository {
//	return &SubjectRepository{db.Table("subjects"), rd}
//}

//Add ... Adds an item
func (r *PurchaseRepository) Add(purchase models.Purchase) string {

	result, err := r.clinet.ConnectPurchases().InsertOne(context.TODO(), purchase)
	if err != nil {
		log.Println(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}

func (r *PurchaseRepository) Find() []*models.Purchase {

	var items []*models.Purchase

	cur, err := r.clinet.ConnectPurchases().Find(context.TODO(), bson.D{})

	if err != nil {
		log.Println(err)
	}

	for cur.Next(context.TODO()) {
		var item models.Purchase
		err := cur.Decode(&item)
		if err != nil {
			log.Println(err)
		}

		items = append(items, &item)

	}

	return items
}
