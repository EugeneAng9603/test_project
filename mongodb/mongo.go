package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBSaver struct {
	collection *mongo.Collection
}

func (s *MongoDBSaver) Save(data string) error {
	_, err := s.collection.InsertOne(context.TODO(), map[string]string{"data": data})
	return err
}

func NewMongoDBSaver(uri, dbName, collName string) (*MongoDBSaver, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	collection := client.Database(dbName).Collection(collName)
	return &MongoDBSaver{collection: collection}, nil
}
