package models

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestCreateDB(t *testing.T) {
	db, err := CreateDB("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	db.Create(&User{BaseModel: BaseModel{Name: "Theo"}, Email: "theo.sun@outlook.com"})

	u := &User{}

	db.Where(&User{BaseModel: BaseModel{Name: "Theo"}}).Take(u)

	if u.ID == 0 {
		t.Fatal("create failed")
	}

}
