package goods

import (
	"encoding/json"
	"strconv"
)

func (s *service) DeleteCategory(reqId string, id string) ([]byte, error) {

	ids, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	category, err := s.db.DeleteCategory(reqId, ids)

	if err != nil {
		s.logger.L.WithField("SERVICE.DeleteCategory", reqId).Error(err)
		return nil, err
	}

	s.logger.L.WithField("SERVICE.DeleteCategory", reqId).Debug("Query row in service layer - ", category)
	SL, err := json.Marshal(category)
	if err != nil {
		s.logger.L.WithField("SERVICE.DeleteCategory", reqId).Error(err)
		return nil, err
	}

	return SL, nil
}
