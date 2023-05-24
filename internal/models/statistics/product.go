package statistics_models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	pb "github.com/MSFT/pkg/services/statistics"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	Uuid  string `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100);not null"`
	Count int    `gorm:"not null"`
	Type  int32  `gorm:"not null"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Uuid = uuid.NewString()
	return
}

func (p Product) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Product) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(source, &p)
}

func (p *Product) ToGRPCModel() *pb.Product {
	return &pb.Product{
		Uuid:        p.Uuid,
		Name:        p.Name,
		Count:       int64(p.Count),
		ProductType: pb.StatisticsProductType(p.Type),
	}
}

type ProductArray []Product

func (pa *ProductArray) Scan(src any) error {
	return pq.GenericArray{A: pa}.Scan(src)
}

func (pa ProductArray) Value() (driver.Value, error) {
	return pq.GenericArray{A: pa}.Value()
}

func (pa ProductArray) ToGRPCModel() []*pb.Product {
	var products []*pb.Product

	for _, product := range pa {
		products = append(products, product.ToGRPCModel())
	}

	return products
}
