package Persistence

import (
	DomainEntities "OrderSubscriber/Domain/Entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
}

func NewDatabase(url string) (*Database, error) {
	cfg := &gorm.Config{}

	db, err := gorm.Open(mysql.Open(url), cfg)
	if err != nil {
		return nil, err
	}

	instance := &Database{
		Gorm: db,
	}

	err = instance.setup()
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (rDb *Database) setup() error {
	return rDb.Gorm.AutoMigrate(
		&DomainEntities.OrderEntity{},
	)
}

func (rDb *Database) Close() error {
	db, err := rDb.Gorm.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}
