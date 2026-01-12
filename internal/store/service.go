package store

import (
	"fmt"
	"log"
	"time"

	"github.com/adi-QTPi/thestral/internal/admin/dto"
	"github.com/adi-QTPi/thestral/internal/config"
	"github.com/adi-QTPi/thestral/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service interface {
	Create(input dto.RouteInput) error
	FindOneRoute(filter *model.Route) (*model.Route, error)
	FindManyRoutes(filter *model.Route) ([]model.Route, error)
}

type service struct {
	cfg *config.Env
	db  *gorm.DB
}

func NewService(cfg *config.Env) (Service, error) {
	var db *gorm.DB

	db, err := gorm.Open(postgres.Open(cfg.DATABASE_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(25) //depends purely on the number of proxy servers
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Migrating database schema")
	if err := db.AutoMigrate(&model.Route{}); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	service := &service{
		cfg: cfg,
		db:  db,
	}

	return service, nil
}
