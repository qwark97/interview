package controller

import (
	"net/http"

	"github.com/qwark97/interview/store/model"
)

type Store interface {
	InsertUser(user model.User) error
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
