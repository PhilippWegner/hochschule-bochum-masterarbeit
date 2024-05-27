package database

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbHost = "localhost"
	dbPort = "27017"
	dbUser = "admin"
	dbPass = "Password"
	dbName = "logs"
)

type Mongodb struct {
	client *mongo.Client
}

type LogEntry struct {
	Name      string    `json:"name" bson:"name"`
	Data      string    `json:"data" bson:"data"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func MongdbConnect() *Mongodb {
	// clientOptions := options.Client().ApplyURI("mongodb://" + dbHost + ":" + dbPort)
	// clientOptions.SetAuth(options.Credential{
	// 	Username: dbUser,
	// 	Password: dbPass,
	// })

	// client, err := mongo.Connect(context.Background(), clientOptions)
	// if err != nil {
	// 	log.Fatal("Error connecting to MongoDB: ", err)
	// }
	// return &Mongodb{client: client}
	return &Mongodb{}
}

func (db *Mongodb) Insert(entry LogEntry) error {
	log.Println("Inserting log entry: ", entry)
	// collection := db.client.Database(dbName).Collection("logs")
	// _, err := collection.InsertOne(context.Background(), LogEntry{
	// 	Name:      entry.Name,
	// 	Data:      entry.Data,
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// })
	// if err != nil {
	// 	log.Println("Error inserting log entry: ", err)
	// 	return err
	// }
	// return err
	return nil
}
