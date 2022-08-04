package Repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golangcodebase/Domain/Entities"
	DomainInterfaces "golangcodebase/Domain/Interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type Suite struct {
	db         *gorm.DB
	mock       sqlmock.Sqlmock
	NameEntity Entities.NameEntity
}

func TestGORMV2(t *testing.T) {
	s := &Suite{}
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}

	if db == nil {
		t.Error("mock db is null")
	}

	if s.mock == nil {
		t.Error("sqlmock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})
	s.db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	if s.db == nil {
		t.Error("gorm db is null")
	}

	defer db.Close()

	var name = &Entities.NameEntity{
		Id:   1,
		Name: "Ehsan",
	}

	t.Run("Rollback on error", func(t *testing.T) {
		unitOfWork := NewUnitOfWork(s.db)
		err = unitOfWork.Do(func(work DomainInterfaces.IUnitOfWork) error {
			_, err = work.NameRepositories().Add(name)
			require.NoError(t, err)

			name.Id = 2
			name.Name = "ali"
			_, err = work.NameRepositories().Add(name)
			require.NoError(t, err)

			return nil
		})

		if assert.ErrorIs(t, err, nil) {
			find, _ := unitOfWork.NameRepositories().Find()
			assert.EqualValues(t, find, nil)
		}
	})

}
