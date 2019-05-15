package models

import (
	"os"

	"github.com/jinzhu/gorm"
)

const (
	// KeyDatabaseType const
	KeyDatabaseType = "DATABASE_TYPE"
	// KeyConnectionString const
	KeyConnectionString = "CONNECTION_STR"
)

// CreateDB instance
func CreateDB() (*gorm.DB, error) {
	dialet := os.Getenv(KeyDatabaseType)
	connStr := os.Getenv(KeyConnectionString)

	if dialet == "" {
		dialet = "sqlite3"
	}
	if connStr == "" {
		connStr = ":memory:"
	}

	db, err := gorm.Open(dialet, connStr)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{}, &Tenant{}, &Solution{})

	return db, err

}
