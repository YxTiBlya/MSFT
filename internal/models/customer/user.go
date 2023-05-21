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

type User struct {
	Uuid        string    `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(50);not null"`
	Office_uuid string    `gorm:"type:varchar(100);not null"`
	Office_name string    `gorm:"type:varchar(100);not null"`
	CreatedAt   time.Time `gorm:"type:timestamp without time zone; not null"`
}

func (p *User) BeforeCreate(tx *gorm.DB) (err error) {
	p.Uuid = uuid.NewString()
	return
}

func (p User) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *User) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(source, &p)
}

type UserArray []User

func (ua *UserArray) Scan(src any) error {
	return pq.GenericArray{A: ua}.Scan(src)
}

func (ua UserArray) Value() (driver.Value, error) {
	return pq.GenericArray{A: ua}.Value()
}
