package models

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (s *NullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	} else {
		return json.Marshal(nil)
	}
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	var val *string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	if val != nil {
		s.Valid = true
		s.String = *val
	} else {
		s.Valid = false
	}

	return nil
}
