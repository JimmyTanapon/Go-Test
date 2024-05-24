package main

import (
	"testing"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setUpTestDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cahe=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&User{})
	return db
}

func TestUser(t *testing.T) {
	db := setUpTestDb()

	t.Run("Successfully add User", func(t *testing.T) {
		err := AddUser(db, "John Doe", "jane.doe@example.com", 30)
		assert.NoError(t, err)
		var user User
		db.First(&user, "email = ?", "jane.doe@example.com")
		assert.Equal(t, "John Doe", user.Fullname)
	})
	t.Run("Fail to add user with existing email ", func(t *testing.T) {
		err := AddUser(db, "John Doe", "jane.doe@example.com", 29)
		assert.EqualError(t, err, "email already exists")
	})
}
