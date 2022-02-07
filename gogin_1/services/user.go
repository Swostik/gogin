package services

import (
	"encoding/json"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

type User struct {
	FirstName    string `json:"firstName`
	LastName     string `json:"lastName"`
	PersonalCode string `json:"personalCode"`
}

var Users []User

type UserService interface {
	//creating user and returning tuple of user an error
	Create(p User) (*User, error)
	Get(code string) (*User, error)
	Update(p User) (*User, error)
	Delete(code string) error
	List() ([]*User, error)
}

type userService struct {
	db *leveldb.DB
}

func NewuserService(db *leveldb.DB) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) Create(p User) (*User, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	err = s.db.Put([]byte(p.PersonalCode), []byte(data), nil)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (s *userService) Get(code string) (*User, error) {
	data, err := s.db.Get([]byte(code), nil)
	if err != nil {
		return nil, err
	}
	var user User
	return &user, json.Unmarshal(data, &user)
}
func (s *userService) Update(p User) (*User, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	err = s.db.Put([]byte(p.PersonalCode), []byte(data), nil)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (s *userService) Delete(code string) error {
	return s.db.Delete([]byte(code), nil)
}
func (s *userService) List() ([]*User, error) {
	iter := s.db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		//key := iter.Key()
		value := iter.Value()
		fmt.Println(value)

	}
	iter.Release()
	return nil, nil

}
