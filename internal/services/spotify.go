package services

import (
	"context"
	"fmt"

	"github.com/jasperspahl/satpl/internal/database"
	"github.com/jasperspahl/satpl/internal/spotify"
	"golang.org/x/oauth2"
)

func getSpotifyClient(ctx context.Context, queries *database.Queries, userID int) (client spotify.SpotifyClient, err error) {
	client, _, err = getSpotifyClientAndUser(ctx, queries, userID)
	return
}

func getSpotifyClientAndUser(ctx context.Context, queries *database.Queries, userID int) (spotify.SpotifyClient, database.User, error) {
	fmt.Println("[SERVICE]: Getting Spotify Client and User")
	user, err := queries.GetUserByID(ctx, int32(userID))
	if err != nil {
		return nil, database.User{}, err
	}
	token := &oauth2.Token{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
	client := spotify.Client(ctx, token)
	return client, user, nil
}
