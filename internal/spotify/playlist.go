package spotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *client) CreatePlaylist(userID string, name string, public bool) (*Playlist, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/users/%s/playlists", userID)
	data := map[string]interface{}{
		"name":   name,
		"public": public,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	fmt.Printf("[SPOTIFY]: creating playlist %s\n", name)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Printf("[SPOTIFY]: returned with %s\n", resp.Status)
	if resp.StatusCode >= 400 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("an error when creating the playlist on spotify")
		}
		fmt.Printf("[SPOTIFY]: %s\n", string(body))
		return nil, fmt.Errorf("spotify: could not create playlist: %s", string(body))
	}
	playlist := new(Playlist)
	err = json.NewDecoder(resp.Body).Decode(playlist)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}
