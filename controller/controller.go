package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/qwark97/interview/fetcher"
	fetcherModel "github.com/qwark97/interview/fetcher/model"
	storeModel "github.com/qwark97/interview/store/model"
)

type Store interface {
	InsertUser(user storeModel.User) error
}

type Fetcher interface {
	Users(ctx context.Context) <-chan fetcherModel.DataBatch
}

type Controller struct {
	storage Store
	fetcher Fetcher
}

func New(storage Store, fetcher fetcher.Fetcher) Controller {
	return Controller{storage: storage, fetcher: fetcher}
}

func (c Controller) Handle(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithCancel(request.Context())
	defer cancel()

	dataCh := c.fetcher.Users(ctx)
	defer drain(dataCh)

	for data := range dataCh {
		if data.Error != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, user := range data.Users {
			storageUser := storeModel.User{
				ID:       user.ID,
				FullName: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
			}
			err := c.storage.InsertUser(storageUser)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

func drain[T any](ch <-chan T) {
	for range ch {
	}
}
