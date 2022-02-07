package services

import (
	"testing"

	"github.com/syndtr/goleveldb/leveldb"
)

func TestCreateUser(t *testing.T) {
	db, err := leveldb.OpenFile("/tmp/example.db", nil)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	service := userService{
		db: db,
	}

	user := User{
		FirstName:    "Hello",
		LastName:     "World",
		PersonalCode: "123",
	}

	_, err = service.Create(user)
	if err != nil {
		t.Errorf("create user fails: %v", err)
	}
}

func TestGetUser(t *testing.T) {
	db, err := leveldb.OpenFile("/tmp/example.db", nil)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	service := userService{
		db: db,
	}

	_, err = service.Get("123")
	if err != nil {
		t.Errorf("create user fails: %v", err)
	}
}
