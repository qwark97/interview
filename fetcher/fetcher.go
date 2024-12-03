package fetcher

import (
	"context"
	"encoding/json"
	"iter"
	"net/http"
	"net/url"

	"github.com/qwark97/interview/fetcher/model"
)

const (
	usersEndpoint = "/v1/users"
)

type Fetcher struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string, client *http.Client) Fetcher {
	return Fetcher{baseURL: baseURL, client: client}
}

func (f Fetcher) FetchUsers(ctx context.Context) (iter.Seq[model.IterData[model.User]], error) {
	currentURL, _ := url.JoinPath(f.baseURL, usersEndpoint)

	return func(yield func(model.IterData[model.User]) bool) {
		for currentURL != "" {
			var iterData model.IterData[model.User]
			apiResponse, err := func() (model.Response, error) {
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, currentURL, nil)
				if err != nil {
					return model.Response{}, err
				}

				resp, err := f.client.Do(req)
				if err != nil {
					return model.Response{}, err
				}

				if resp.StatusCode == http.StatusTooManyRequests {
					return model.Response{}, nil
				}

				var data model.Response
				if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
					return model.Response{}, err
				}

				resp.Body.Close()
				return data, nil
			}()
			if err != nil {
				iterData.Error = err
				yield(iterData)
				return
			}

			iterData.Data = apiResponse.Users
			if !yield(iterData) {
				return
			}

			currentURL = apiResponse.NextLink
		}
	}, nil
}
