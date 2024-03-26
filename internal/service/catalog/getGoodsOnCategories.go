package catalog

import "encoding/json"

func (s *catalog) GoodsOnCateory(reqId string, category string) ([]byte, error) {
	categories, err := s.db.GoodsOnCateory(reqId, category)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	s.logger.L.WithField("SERVICE.CreateCategory", reqId).Debug("Query row in catalog layer - ", category)
	SL, err := json.Marshal(categories)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	return SL, nil
}
