package hevy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// apiTransport is a custom transport for API requests
type apiTransport struct {
	apiKey string
	agent  string
	base   http.RoundTripper
}

// roundTrip is a custom roundtripper that adds the necessary request fields
// for API requests
func (t apiTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", t.agent)
	req.Header.Add("api-key", t.apiKey)

	base := t.base
	if t.base == nil {
		base = http.DefaultTransport
	}

	return base.RoundTrip(req)
}

// paginated is a base class wrapper for working with paginated results
type paginatedResults struct {
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

// Construct a URL for querying the API.
// if `page` is not 0, append the paginated query strings
// to the request.
func (c Client) constructURL(path string, query map[string]string) string {
	base := fmt.Sprintf("%s/%s/%s", c.ApiURL, c.ApiVersion, path)

	queryString := url.Values{}
	if len(query) > 0 {
		for k, v := range query {
			queryString.Add(k, v)
		}
	}

	return fmt.Sprintf("%s?%s", base, queryString.Encode())
}

// request a single API endpoint.  Data is writen to the pointer
// given in the resp var.
func (c Client) get(url string, resp any) error {
	data, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer data.Body.Close()

	body, err := io.ReadAll(data.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, resp)
	if err != nil {
		return err
	}

	return nil
}
