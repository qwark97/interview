package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qwark97/interview/fetcher/model"
)

type Fetcher struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string, client *http.Client) Fetcher {
	return Fetcher{
		baseURL: baseURL,
		client:  client,
	}
}

func (f Fetcher) Users(ctx context.Context) <-chan model.DataBatch {
	dataCh := make(chan model.DataBatch)

	go func() {
		defer close(dataCh)

		url := fmt.Sprintf("%s/users?page=1", f.baseURL)
		for url != "" {
			select {
			case <-ctx.Done():
				return
			default:
			}
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				dataCh <- model.DataBatch{
					Users: nil,
					Error: err,
				}
				return
			}

			resp, err := f.client.Do(req)
			if err != nil {
				dataCh <- model.DataBatch{
					Users: nil,
					Error: err,
				}
				return
			}

			if resp.StatusCode == http.StatusTooManyRequests {
				continue
			}

			var response model.Response
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				dataCh <- model.DataBatch{
					Users: nil,
					Error: err,
				}
				return
			}

			dataCh <- model.DataBatch{
				Users: response.Users,
				Error: nil,
			}
			url = response.NextLink
		}
	}()

	return dataCh
}
