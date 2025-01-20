package services

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jasperspahl/satpl/internal/database"
	"github.com/jasperspahl/satpl/internal/spotify"
	"golang.org/x/oauth2"
)

type UserService interface {
	LoginUser(ctx context.Context, userProfile spotify.UserProfile, token *oauth2.Token) (int, error)
}

type userService struct {
	queries *database.Queries
}

func NewUserService(q *database.Queries) UserService {
	return &userService{
		queries: q,
	}
}

func (s *userService) LoginUser(ctx context.Context, userProfile spotify.UserProfile, token *oauth2.Token) (int, error) {
	user, err := s.queries.GetUserBySpotifyID(ctx, userProfile.ID)
	if err == sql.ErrNoRows {
		return s.createNewUser(ctx, userProfile, token)
	} else if err != nil {
		return -1, err
	}

	err = s.queries.UpdateTokens(ctx, database.UpdateTokensParams{
		ID:           user.ID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	})
	if err != nil {
		return -1, err
	}

	return int(user.ID), nil
}

func (s *userService) createNewUser(ctx context.Context, userProfile spotify.UserProfile, token *oauth2.Token) (int, error) {
	user, err := s.queries.CreateUser(ctx, database.CreateUserParams{
		SpotifyID: userProfile.ID,
		DisplayName: pgtype.Text{
			String: userProfile.DisplayName,
			Valid:  true,
		},
		Email: pgtype.Text{
			String: userProfile.Email,
			Valid:  true,
		},
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	})
	if err != nil {
		return -1, err
	}
	return int(user.ID), nil
}
