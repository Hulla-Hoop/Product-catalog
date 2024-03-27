package autification

import "time"

// Чистит базу от устаревших сессий
func (s *jWT) ClearSession() {
	for {
		time.Sleep(time.Minute * 5)
		s.db.DeleteOld()
	}
}
