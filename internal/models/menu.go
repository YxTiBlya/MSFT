package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	Uuid            string       `gorm:"primaryKey"`
	OnDate          string       `gorm:"type:varchar(100);not null"`
	OpeningRecordAt string       `gorm:"type:varchar(100);not null"`
	ClosingRecordAt string       `gorm:"type:varchar(100);not null"`
	Salads          ProductArray `gorm:"type:jsonb[]"`
	Garnishes       ProductArray `gorm:"type:jsonb[]"`
	Meats           ProductArray `gorm:"type:jsonb[]"`
	Soups           ProductArray `gorm:"type:jsonb[]"`
	Drinks          ProductArray `gorm:"type:jsonb[]"`
	Desserts        ProductArray `gorm:"type:jsonb[]"`
	CreatedAt       string       `gorm:"type:varchar(100);not null"`
}

type CreateMenuRequest struct {
	OnDate          time.Time `json:"on_date"`
	OpeningRecordAt time.Time `json:"opening_record_at"`
	ClosingRecordAt time.Time `json:"closing_record_at"`
	Salads          []string  `json:"salads"`
	Garnishes       []string  `json:"garnishes"`
	Meats           []string  `json:"meats"`
	Soups           []string  `json:"soups"`
	Drinks          []string  `json:"drinks"`
	Desserts        []string  `json:"desserts"`
}

func (p *Menu) BeforeCreate(tx *gorm.DB) (err error) {
	p.Uuid = uuid.NewString()
	return
}

func (m Menu) Value() (driver.Value, error) {
	return json.Marshal(m)
}
