package services

import (
	"context"
	"fmt"

	"github.com/jasperspahl/satpl/internal/database"
	"github.com/jasperspahl/satpl/internal/models"
)

type SearchService interface {
	Search(ctx context.Context, userID int, q string) ([]models.Artist, error)
}

type searchService struct {
	queries *database.Queries
}

func NewSearchService(q *database.Queries) SearchService {
	return &searchService{q}
}

func (s *searchService) Search(ctx context.Context, userID int, q string) ([]models.Artist, error) {
	fmt.Printf("[SEARCH]: Searching for %s\n", q)
	var artists []models.Artist
	client, err := getSpotifyClient(ctx, s.queries, userID)
	if err != nil {
		return artists, err
	}
	result, err := client.Search(q)
	if err != nil {
		return artists, err
	}
	for _, artist := range result {
		artists = append(artists, models.Artist{
			ID:          artist.ID,
			Name:        artist.Name,
			Genres:      artist.Genres,
			Popularity:  artist.Popularity,
			ExternalUrl: artist.ExternalUrls.Spotify,
		})
		if len(artist.Images) > 0 {
			artists[len(artists)-1].Image = artist.Images[0].Url
		}
	}
	return artists, nil
}
