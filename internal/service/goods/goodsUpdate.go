package goods

import (
	"encoding/json"
	"strconv"
)

func (s *service) UpdateGoods(reqId string, id string, name string) ([]byte, error) {

	ids, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	category, err := s.db.UpdateGoods(reqId, ids, name)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	s.logger.L.WithField("SERVICE.CreateCategory", reqId).Debug("Query row in service layer - ", category)
	SL, err := json.Marshal(category)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	return SL, nil
}
