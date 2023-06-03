package client

import (
	"context"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"os"
)

type API interface {
	FindPlanetFilmAppearances(ctx context.Context, name string) (int8, error)
}

type Client struct {
	Client *http.Client
}

const (
	path        = "/api/planets"
	searchParam = "search"
)

func (c *Client) FindPlanetFilmAppearances(ctx context.Context, name string) (int8, error) {
	host := os.Getenv("SW_HOST")
	req, err := http.NewRequest(http.MethodGet, host+path, nil)

	q := req.URL.Query()
	q.Add(searchParam, name)

	req.URL.RawQuery = q.Encode()
	response, err := c.Client.Do(req)

	if err != nil || response.StatusCode != http.StatusOK {
		return 0, err
	}

	body, _ := io.ReadAll(response.Body)

	var results struct {
		Results []struct {
			Films []string
			Name  string
		}
	}

	err = sonic.Unmarshal(body, &results)
	if err != nil || len(results.Results) == 0 {
		return 0, err
	}

	return int8(len(results.Results[0].Films)), nil
}
