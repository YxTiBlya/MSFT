package customer_models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Office struct {
	Uuid      string    `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(50);not null"`
	Address   string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"not null"`
}

func (p *Office) BeforeCreate(tx *gorm.DB) (err error) {
	p.Uuid = uuid.NewString()
	return
}

func (p Office) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Office) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(source, &p)
}

type OfficeArray []Office

func (oa *OfficeArray) Scan(src any) error {
	return pq.GenericArray{A: oa}.Scan(src)
}

func (oa OfficeArray) Value() (driver.Value, error) {
	return pq.GenericArray{A: oa}.Value()
}
