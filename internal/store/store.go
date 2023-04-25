package store

import (
	"fmt"

	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnToDB(c *cfg.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v", c.Postgres_username, c.Postgres_password, c.Postgres_host, c.Postgres_port, c.Postgres_dbname, c.Postgres_ssl)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(
		&models.Products{},
	)

	return DB, nil
}
