package spotify

import "encoding/json"

func (c *client) GetCurrentUser() (UserProfile, error) {
	resp, err := c.httpClient.Get("https://api.spotify.com/v1/me")
	if err != nil {
		return UserProfile{}, err
	}
	defer resp.Body.Close()
	var user UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return UserProfile{}, err
	}
	return user, nil
}
