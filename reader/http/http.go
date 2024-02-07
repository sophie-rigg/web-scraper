package http

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type Client struct {
	http.Client
}

func New() *Client {
	return &Client{
		Client: http.Client{},
	}
}

//TODO: Add a sever that spins up a test page to test this function with

// Read takes a URL and returns a html.Node
func (c *Client) Read(url string) (*html.Node, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var tries int

	for {
		response, err := c.Do(request)
		if err != nil {
			return nil, fmt.Errorf("failed to make request: %w", err)
		}

		// if the status code is not 200 then we will retry
		if response.StatusCode != http.StatusOK {
			tries++
			// give 3 tries before failing
			if tries > 3 {
				return nil, fmt.Errorf("failed to make request: %w", err)
			}
			continue
		}
		defer response.Body.Close()

		// parse the response into a html node
		node, err := html.Parse(response.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
		}
		return node, nil
	}
}
