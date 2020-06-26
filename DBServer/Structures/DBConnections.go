package Structures

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DBConnections struct {
	MongoClient     *mongo.Client
	MemeIndexerDB   *mongo.Database
	UsersCollection *mongo.Collection
	MemesCollection *mongo.Collection
}

func (DB *DBConnections) InitializeDBConnections() {
	DB.MongoClient = initializeMongoConnection()
	DB.MemeIndexerDB = connectToMemeIndexerDB(DB.MongoClient)

	const UsersCollection string = "Users"
	const MemesCollection string = "Memes"

	DB.UsersCollection = connectToCollection(UsersCollection, DB.MemeIndexerDB)
	DB.MemesCollection = connectToCollection(MemesCollection, DB.MemeIndexerDB)
}

func initializeMongoConnection() (mongoClient *mongo.Client) {
	var err error
	const DBAddress string = "mongodb://localhost:27017"

	c, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err = mongo.Connect(c, options.Client().ApplyURI(DBAddress))
	if err != nil {
		log.Panic(err)
	}

	return
}

func connectToMemeIndexerDB(mongoClient *mongo.Client) (memeIndexerDB *mongo.Database) {
	const DBName string = "MemeIndexer"

	memeIndexerDB = mongoClient.Database(DBName)
	return
}

func connectToCollection(collectionName string, memeIndexerDB *mongo.Database) (collectionConnection *mongo.Collection) {
	collectionConnection = memeIndexerDB.Collection(collectionName)
	return
}
