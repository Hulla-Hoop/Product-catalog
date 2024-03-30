package catalog

import (
	"encoding/json"
)

func (s *catalog) CreateGoods(reqId string, name string, category string) ([]byte, error) {

	goods, err := s.db.CreateGoods(reqId, name)

	if err != nil {
		s.logger.L.WithField("SERVICE.CreateGoods", reqId).Error(err)
		return nil, err
	}

	cat, err := s.db.CreateCategory(reqId, category)
	if err != nil {
		return nil, err
	}

	product, err := s.db.CreateRelation(reqId, goods.ID, cat.ID)

	if err != nil {
		s.logger.L.WithField("SERVICE.CreateGoods", reqId).Error(err)
		return nil, err
	}

	s.logger.L.WithField("SERVICE.CreateGoods", reqId).Debug("Query row in catalog layer - ", product)

	SL, err := json.Marshal(product)

	if err != nil {
		s.logger.L.WithField("SERVICE.CreateGoods", reqId).Error(err)
		return nil, err
	}

	return SL, nil
}
