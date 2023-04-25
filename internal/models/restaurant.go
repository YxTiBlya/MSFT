package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Products struct {
	//ID          string  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Id          string  `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text;not null"`
	Type        int32   `gorm:"not null"`
	Weight      int32   `gorm:"not null"`
	Price       float64 `gorm:"type:double precision"`
	CreatedAt   string  `gorm:"type:varchar(100);not null"`
}

func (p *Products) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id = uuid.NewString()
	return
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
