package Persistence

import (
	"errors"
	GoCache "github.com/patrickmn/go-cache"
	DomainEntities "golangcodebase/Domain/Entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Instance struct {
	Gorm        *gorm.DB
	MemoryCache *GoCache.Cache
}

func New(url string, log bool) (*Instance, error) {
	var logMode logger.LogLevel
	if log {
		logMode = logger.Info
	}

	cfg := &gorm.Config{Logger: logger.Default.LogMode(logMode), SkipDefaultTransaction: true}

	db, err := gorm.Open(postgres.Open(url), cfg)
	if err != nil {
		return nil, err
	}
	i := &Instance{
		Gorm: db,
	}

	return i, nil
}

func Setup(db *gorm.DB) error {
	return db.AutoMigrate(
		&DomainEntities.NameEntity{},
		&DomainEntities.Name1Entity{},
	)
}

func (i *Instance) MemorySet(key string, data []byte, exp time.Duration) error {
	i.MemoryCache.Set(key, data, exp)
	return nil
}

func (i *Instance) MemoryGet(key string) ([]byte, error) {
	res, ok := i.MemoryCache.Get(key)
	if !ok {
		return nil, errors.New("not found")
	}
	return res.([]byte), nil
}
