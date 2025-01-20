package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	resp, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	playlist := new(Playlist)
	err = json.NewDecoder(resp.Body).Decode(playlist)
	if err != nil {
		return nil, err
	}
	return playlist, nil
}
