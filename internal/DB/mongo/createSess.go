package mongo

import (
	"context"
	"testinhousead/internal/model"
)

func (m *Mongo) CreateSess(reqId string, session *model.Session) error {
	_, err := m.collection.InsertOne(context.TODO(), session)
	if err != nil {
		return err
	}
	return nil

}
