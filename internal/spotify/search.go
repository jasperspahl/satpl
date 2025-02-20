package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
)

func (c *client) Search(q string) ([]Artist, error) {
	fmt.Printf("[SPOTIFY]: searching for %s\n", q)
	url := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist", url.QueryEscape(q))
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("error while searching on spotify")
		}
		fmt.Printf("[SPOTIFY]: %s\n", string(body))
		return nil, fmt.Errorf("spotify: error while searching: %s", string(body))
	}
	defer resp.Body.Close()
	results := new(SearchResults)
	err = json.NewDecoder(resp.Body).Decode(results)
	if err != nil {
		return nil, err
	}
	return results.Artists.Items, nil
}
