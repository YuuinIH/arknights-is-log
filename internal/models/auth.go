package models

import (
	"context"

	"github.com/google/uuid"
	u "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Login struct {
	UUID string `json:"uuid" form:"uuid" example:"00000000-0000-0000-0000-000000000000"`
}

type NewToken struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}

func NewUUID() (u.UUID, error) {
	newuuid, err := u.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}
	_, err = db.Database("is_log").Collection("account").InsertOne(context.TODO(), bson.D{
		{"uuid", newuuid.String()},
	})
	return newuuid, err
}

func CheckUUID(uuid u.UUID) (bool, error) {
	filter := bson.D{{"uuid", uuid.String()}}
	result := new(bson.M)
	err := db.Database("is_log").Collection("account").FindOne(context.TODO(), filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func BindEmailToUUID(nuuid uuid.UUID, email string) error {
	filter := bson.D{{"uuid", nuuid}}
	_, err := db.Database("is_log").Collection("account").ReplaceOne(context.TODO(), filter,
		bson.D{{"$set", bson.D{{"email", email}}}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return err
		}
		return err
	}
	return nil
}
