package main

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an Error '%s' was not expect when opening a stub data connection ", err)

	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}
	defer db.Close()
	t.Run("Add User successfully", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs("jane.doe@example.com").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO \"users\" (.+)$").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err := AddUser(gormDB, "John Doe", "jane.doe@example.com", 30)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())

	})
	t.Run("fail to add user with existing email", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
			WithArgs("jane.doe@example.com").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		err := AddUser(gormDB, "John Doe", "jane.doe@example.com", 30)
		assert.EqualError(t, err, "email already exists")
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
