package mongo

import (
	"context"
	"testinhousead/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// проверяет наличие сессии в базе по токену и возвращает сессию в случае успеха
func (m *Mongo) ChekSess(reqId string, token string) (*model.Session, error) {
	var sess model.Session
	filter := bson.D{primitive.E{Key: "bcryptTocken", Value: token}}
	res := m.collection.FindOne(context.TODO(), filter)

	err := res.Decode(&sess)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}
