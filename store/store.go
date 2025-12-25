package store

import (
	"fmt"

	"github.com/qwark97/interview/store/model"
)

type Store struct {
}

func New() Store {
	return Store{}
}

func (s Store) InsertUser(user model.User) error {
	/*
		...
	*/
	fmt.Println("Inserting user:", user.FullName)
	return nil
}
