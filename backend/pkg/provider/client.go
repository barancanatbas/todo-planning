package provider

import (
	"io"
	"net/http"
	"net/url"
	"task-manager/internal/model/entity"
)

type client struct {
}

func NewProviderClient() *client {
	return &client{}
}

func (c *client) Get(url *url.URL, providerModel ITask) ([]entity.Task, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result, err := providerModel.ConvertAll(body)
	return result, err
}
