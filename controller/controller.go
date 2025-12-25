package controller

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"net/http"
	"time"

	apiModel "github.com/qwark97/interview/fetcher/model"
	dbModel "github.com/qwark97/interview/store/model"
)

type Store interface {
	InsertUser(user dbModel.User) error
}

type Fetcher interface {
	FetchUsers(ctx context.Context) (iter.Seq[apiModel.IterData[apiModel.User]], error)
}

type Controller struct {
	storage Store
	fetcher Fetcher
}

func New(storage Store, fetcher Fetcher) Controller {
	return Controller{storage: storage, fetcher: fetcher}
}

func (c Controller) Handle(writer http.ResponseWriter, request *http.Request) {
	/* TODO:
	- Fetch users from vendor API
	- Insert new users into database
	- In case of processing failure (unrecoverable), return 500
	*/
	ctx, cancel := context.WithTimeout(request.Context(), time.Minute)
	defer cancel()

	iterator, err := c.fetcher.FetchUsers(ctx)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	for user := range iterator {
		if user.Error != nil {
			http.Error(writer, user.Error.Error(), http.StatusInternalServerError)
			return
		}

		for _, user := range user.Data {
			dbUser := dbModel.User{
				ID:       user.ID,
				FullName: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			}
			err := c.storage.InsertUser(dbUser)
			if errors.Is(err, dbModel.ErrDuplicate) {
				continue
			} else if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
