package restaurant_models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/lib/pq"
)

type Orders struct {
	TotalOrders          OrderArray         `gorm:"type:jsonb[]"`
	TotalOrdersByCompany OrderByOfficeArray `gorm:"type:jsonb[]"`
}

func (o Orders) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *Orders) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(source, &o)
}

// -----

type Order struct {
	ProductUuid string `gorm:"type:varchar(100);not null"`
	ProductName string `gorm:"type:varchar(100);not null"`
	Count       int64  `gorm:"not null"`
}

func (o Order) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *Order) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(source, &o)
}

type OrderArray []Order

func (oa *OrderArray) Scan(src any) error {
	return pq.GenericArray{A: oa}.Scan(src)
}

func (oa OrderArray) Value() (driver.Value, error) {
	return pq.GenericArray{A: oa}.Value()
}

// -----

type OrderByOffice struct {
	CompanyUuid   string     `gorm:"type:varchar(100);not null"`
	OfficeName    string     `gorm:"type:varchar(50);not null"`
	OfficeAddress string     `gorm:"type:varchar(100);not null"`
	Result        OrderArray `gorm:"type:jsonb[]"`
}

func (o OrderByOffice) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *OrderByOffice) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(source, &o)
}

type OrderByOfficeArray []OrderByOffice

func (oboa *OrderByOfficeArray) Scan(src any) error {
	return pq.GenericArray{A: oboa}.Scan(src)
}

func (oboa OrderByOfficeArray) Value() (driver.Value, error) {
	return pq.GenericArray{A: oboa}.Value()
}
