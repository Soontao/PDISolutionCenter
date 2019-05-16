package models

import (
	"github.com/jinzhu/gorm"
)

const (
	// KeyDatabaseType const
	KeyDatabaseType = "DATABASE_TYPE"
	// KeyConnectionString const
	KeyConnectionString = "CONNECTION_STR"
)

// CreateDB instance
func CreateDB(dialect, connStr string) (*gorm.DB, error) {

	if dialect == "" {
		dialect = "sqlite3"
	}

	if connStr == "" {
		connStr = ":memory:"
	}

	db, err := gorm.Open(dialect, connStr)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&User{}, &Tenant{}, &Solution{}, &JobRunLog{}, &Schedule{})

	return db, err

}
