package goods

import "encoding/json"

func (s *service) GoodsOnCateory(reqId string, category string) ([]byte, error) {
	categories, err := s.db.GoodsOnCateory(reqId, category)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	s.logger.L.WithField("SERVICE.CreateCategory", reqId).Debug("Query row in service layer - ", category)
	SL, err := json.Marshal(categories)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	return SL, nil
}
