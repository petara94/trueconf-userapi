package models

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"time"
)

type User struct {
	Model
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UserList map[int]User

type UserStore struct {
	ModelStore
	StorePath string   `json:"-"`
	Increment int      `json:"increment"`
	List      UserList `json:"list"`
}

func NewUserStore(storePath string) *UserStore {
	store := UserStore{StorePath: storePath}
	err := store.readStore()
	if err != nil {
		log.Fatal(err)
	}
	return &store
}

func (u *UserStore) Create(user User) (User, error) {
	u.Lock()
	defer u.Unlock()

	u.Increment++

	user.ID = u.Increment
	user.CreatedAt = time.Now().Round(time.Second)
	user.UpdatedAt = user.CreatedAt

	u.List[u.Increment] = user

	err := u.saveStore()
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *UserStore) Update(id int, item User) error {
	u.Lock()
	defer u.Unlock()

	user, ok := u.List[id]
	if !ok {
		return errors.New(USER_NOT_FOUND)
	}
	user.UpdatedAt = time.Now().Round(time.Second)

	// protecting no-modify data
	item.Model = user.Model
	item.Email = user.Email

	user = item

	u.List[id] = user

	err := u.saveStore()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStore) Delete(id int) error {
	u.Lock()
	defer u.Unlock()

	beforeSize := len(u.List)
	delete(u.List, id)

	if beforeSize == len(u.List) {
		return nil
	}

	err := u.saveStore()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStore) Get(id int) (User, error) {
	u.RLock()
	defer u.RUnlock()

	user, ok := u.List[id]
	if !ok {
		return User{}, errors.New(USER_NOT_FOUND)
	}

	return user, nil
}

func (u *UserStore) GetAll() UserList {
	u.RLock()
	defer u.RUnlock()

	return u.List
}

func (u *UserStore) readStore() error {
	u.Lock()
	defer u.Unlock()

	f, err := ioutil.ReadFile(u.StorePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(f, u)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserStore) saveStore() error {
	b, err := json.Marshal(u)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(u.StorePath, b, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
