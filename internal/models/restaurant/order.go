package restaurant_models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	pb "github.com/MSFT/pkg/services/restaurant"
	"github.com/lib/pq"
)

type Orders struct {
	Id                   uint               `gorm:"primary_key"`
	TotalOrders          OrderArray         `gorm:"type:jsonb[];not null"`
	TotalOrdersByCompany OrderByOfficeArray `gorm:"type:jsonb[];not null"`
	CreatedAt            time.Time          `gorm:"type:timestamp without time zone;not null"`
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

func (o *Order) ToGRPCModel() *pb.Order {
	return &pb.Order{
		ProductId:   o.ProductUuid,
		ProductName: o.ProductName,
		Count:       o.Count,
	}
}

type OrderArray []Order

func (oa *OrderArray) Scan(src any) error {
	return pq.GenericArray{A: oa}.Scan(src)
}

func (oa OrderArray) Value() (driver.Value, error) {
	return pq.GenericArray{A: oa}.Value()
}

func (oa OrderArray) ToGRPCModel() []*pb.Order {
	var orders []*pb.Order

	for _, order := range oa {
		orders = append(orders, order.ToGRPCModel())
	}

	return orders
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

func (o *OrderByOffice) ToGRPCModel() *pb.OrdersByOffice {
	return &pb.OrdersByOffice{
		CompanyId:     o.CompanyUuid,
		OfficeName:    o.OfficeName,
		OfficeAddress: o.OfficeAddress,
		Result:        o.Result.ToGRPCModel(),
	}
}

type OrderByOfficeArray []OrderByOffice

func (oboa *OrderByOfficeArray) Scan(src any) error {
	return pq.GenericArray{A: oboa}.Scan(src)
}

func (oboa OrderByOfficeArray) Value() (driver.Value, error) {
	return pq.GenericArray{A: oboa}.Value()
}

func (oboa OrderByOfficeArray) ToGRPCModel() []*pb.OrdersByOffice {
	var orders []*pb.OrdersByOffice

	for _, order := range oboa {
		orders = append(orders, order.ToGRPCModel())
	}

	return orders
}
