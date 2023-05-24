package statistics_models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Statistics struct {
	Id          uint         `gorm:"primary_key"`
	TopProducts ProductArray `gorm:"type:jsonb[];not null"`
	Profit      float64      `gorm:"not null"`
	CreatedAt   time.Time    `gorm:"type:timestamp without time zone;not null"`
}

func (s Statistics) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Statistics) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(source, &s)
}
