package Persistence

import (
	DomainEntities "OrderSubscriber/Domain/Entities"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
}

func NewDatabase(url string) Database {
	cfg := &gorm.Config{}

	db, err := gorm.Open(mysql.Open(url), cfg)
	if err != nil {
		log.Fatalln(err)
	}

	instance := Database{
		Gorm: db,
	}

	err = instance.setup()
	if err != nil {
		log.Fatalln(err)
	}

	return instance
}

func (rDb Database) setup() error {
	return rDb.Gorm.AutoMigrate(
		&DomainEntities.OrderEntity{},
	)
}

func (rDb Database) Close() {
	db, err := rDb.Gorm.DB()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
