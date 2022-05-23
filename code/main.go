package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ahloul/loyalty-reports/infrastructure/config"
	"github.com/ahloul/loyalty-reports/models"
	"github.com/ahloul/loyalty-reports/purchase/repository"
	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
	"log"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
func main() {

	//m := db.MongoC()
	//m.ConnectPurchases()
	//d:=repository.PurchaseRepository{}

	viper := config.NewViper()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	InitializeConfig(router)
	InitializeRouts(router)
	router.Run(viper.Server.Port)

	// initialize viper object

	// to consume messages
	// get kafka reader using environment variables.
	kafkaURL := viper.KAFKA.Url
	topic := "events"
	groupID := "my-group-id"
	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start consuming ... !!")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		type mData struct {
			key string
		}
		var data map[string]mData

		json.Unmarshal(msg.Value, &data)
		fmt.Println("Inser: ", data)

		if string(msg.Key) == "purchased_created" {
			var m models.Purchase
			json.Unmarshal(msg.Value, &m)

			d := repository.PurchaseRepository{}
			d.Add(m)
		}

	}

	return

}
