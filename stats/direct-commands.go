package stats

import (
	"context"
	"encoding/json"
	"github.com/ArnobKumarSaha/mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func direct(db *mongo.Database, mongoCommand string) {
	cmd := bson.D{{mongoCommand, 1}}
	var result bson.M
	err := db.RunCommand(context.TODO(), cmd).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	indentedData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	utils.WriteFile(dir, mongoCommand, indentedData)
}
