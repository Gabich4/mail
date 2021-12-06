package mongorepository

import (
	"context"
	"log"
	"mailsender/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DatabaseName           = "mailSystemDb"
	MessagesCollectionName = "messages"

	WaitingStatus = "waiting"
	SuccessStatus = "success"
	FailStatus    = "fail"
)

// CreateMessage inserts msg into collection.
func (m *MongoClient) CreateMessage(msg models.Message) (string, error) {
	ctx := context.Background()

	mCollection := m.Database(DatabaseName).Collection(MessagesCollectionName)
	res, err := mCollection.InsertOne(ctx, msg)
	if err != nil {
		return "", err
	}
	log.Printf("Message: %s inserted in messages collection with ID: %s", msg, res.InsertedID)
	return primitive.ObjectID.Hex(res.InsertedID.(primitive.ObjectID)), nil
}

// UpdateMessage updates message in the collection
func (m *MongoClient) UpdateMessage(messageId string, status string) error {
	ctx := context.Background()
	mCollection := m.Database(DatabaseName).Collection(MessagesCollectionName)
	update := bson.D{
		{"$set", bson.D{{"status", status}}},
	}

	id, err := primitive.ObjectIDFromHex(messageId)
	if err != nil {
		return err
	}
	res, err := mCollection.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}

	log.Printf("Message with id: %s updated with status: %s, modifiedCount: %d", id, status, res.ModifiedCount)
	return nil
}
