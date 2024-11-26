package controller

import (
	"net/http"

	dbModel "example.com/store/model"
)

type Store interface {
	InsertUser(user dbModel.User) error
}

type Controller struct {
	storage Store
}

func New(storage Store) Controller {
	return Controller{storage: storage}
}

func (c Controller) Handle(writer http.ResponseWriter, request *http.Request) {
	/* TODO:
	- Fetch users from vendor API
	- Insert new users into database
	- In case of processing failure (unrecoverable), return 500
	*/
}
