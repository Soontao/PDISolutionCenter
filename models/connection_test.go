package models

import (
	"testing"
)

func TestCreateDB(t *testing.T) {
	got, err := CreateDB()
	if err != nil {
		t.Fatal(err)
	}
	if !got.NewRecord(&User{BaseModel: BaseModel{Name: "Theo"}}) {
		t.Fatal("Craete failed")
	}

}
