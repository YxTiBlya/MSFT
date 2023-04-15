package models

import (
	"time"
)

type Restaurant struct {
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
