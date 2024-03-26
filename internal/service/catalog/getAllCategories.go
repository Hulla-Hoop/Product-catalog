package catalog

import "encoding/json"

func (s *catalog) AllCategories(reqId string) ([]byte, error) {
	category, err := s.db.AllCategories(reqId)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	s.logger.L.WithField("SERVICE.CreateCategory", reqId).Debug("Query row in catalog layer - ", category)
	SL, err := json.Marshal(category)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	return SL, nil
}
