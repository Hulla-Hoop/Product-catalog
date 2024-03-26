package catalog

import (
	"encoding/json"
	"strconv"
)

func (s catalog) UpdateCategory(reqId string, id string, name string) ([]byte, error) {
	idI, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	category, err := s.db.UpdateCategory(reqId, idI, name)
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
