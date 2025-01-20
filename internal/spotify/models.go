package spotify

type UserProfile struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type Playlist struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Public bool   `json:"public"`
	Tracks struct {
		Href     string          `json:"href"`
		Limit    int             `json:"limit"`
		Offset   int             `json:"offset"`
		Total    int             `json:"total"`
		Next     string          `json:"next,omitempty"`
		Previous string          `json:"previous,omitempty"`
		Items    []PlaylistTrack `json:"items"`
	} `json:"tracks"`
}

type PlaylistTrack struct {
	IsLocal bool  `json:"is_local"`
	Track   Track `json:"track"`
}

type Track struct {
	ID  string `json:"ID"`
	Uri string `json:"uri"`
}
