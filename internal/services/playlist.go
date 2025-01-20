package services

import (
	"context"

	"github.com/jasperspahl/satpl/internal/database"
	"github.com/jasperspahl/satpl/internal/spotify"
	"golang.org/x/oauth2"
)

type PlaylistService interface{}

type playlistService struct {
	queries *database.Queries
}

func NewPlaylistService(q *database.Queries) PlaylistService {
	return &playlistService{
		queries: q,
	}
}

func (s *playlistService) CreatePlaylist(ctx context.Context, userID int, name string, public bool) error {
	user, err := s.queries.GetUserByID(ctx, int32(userID))
	if err != nil {
		return err
	}
	token := &oauth2.Token{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
	client := spotify.Client(ctx, token)
	playlist, err := client.CreatePlaylist(user.SpotifyID, name, public)
	if err != nil {
		return err
	}
	_, err = s.queries.CreatePlaylist(ctx, database.CreatePlaylistParams{
		UserID:    user.ID,
		SpotifyID: playlist.ID,
		Name:      playlist.Name,
	})
	if err != nil {
		return err
	}

	return nil
}
