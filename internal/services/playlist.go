package services

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jasperspahl/satpl/internal/database"
	"github.com/jasperspahl/satpl/internal/models"
)

type PlaylistService interface {
	CreatePlaylist(ctx context.Context, userID int, name string, public bool) (models.Playlist, error)
	GetPlaylist(ctx context.Context, userID int) ([]models.Playlist, error)
}

type playlistService struct {
	queries *database.Queries
}

func NewPlaylistService(q *database.Queries) PlaylistService {
	return &playlistService{
		queries: q,
	}
}

func (s *playlistService) CreatePlaylist(ctx context.Context, userID int, name string, public bool) (models.Playlist, error) {
	client, user, err := getSpotifyClientAndUser(ctx, s.queries, userID)
	if err != nil {
		return models.Playlist{}, err
	}
	playlist, err := client.CreatePlaylist(user.SpotifyID, name, public)
	if err != nil {
		return models.Playlist{}, err
	}
	dbPlaylist, err := s.queries.CreatePlaylist(ctx, database.CreatePlaylistParams{
		UserID:    user.ID,
		SpotifyID: playlist.ID,
		Name:      playlist.Name,
	})
	if err != nil {
		return models.Playlist{}, err
	}

	return models.Playlist{ID: int(dbPlaylist.ID), SpotifyID: dbPlaylist.SpotifyID, Name: playlist.Name}, nil
}

func (s *playlistService) GetPlaylist(ctx context.Context, userID int) ([]models.Playlist, error) {
	playlists, err := s.queries.GetPlaylistsByUserID(ctx, int32(userID))
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}
	result := make([]models.Playlist, len(playlists))
	for i, playlist := range playlists {
		result[i].ID = int(playlist.ID)
		result[i].Name = playlist.Name
		result[i].SpotifyID = playlist.SpotifyID
	}
	return result, nil
}
