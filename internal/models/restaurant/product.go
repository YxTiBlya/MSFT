package restaurant_models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/MSFT/internal/timestamp"
	pb "github.com/MSFT/pkg/services/restaurant"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	Uuid        string    `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text;not null"`
	Type        int32     `gorm:"not null"`
	Weight      int32     `gorm:"not null"`
	Price       float64   `gorm:"type:double precision"`
	CreatedAt   time.Time `gorm:"type:timestamp without time zone; not null"`
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
		Description: p.Description,
		Type:        pb.ProductType(p.Type),
		Weight:      p.Weight,
		Price:       p.Price,
		CreatedAt:   timestamp.ToTimestamppb(p.CreatedAt),
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
