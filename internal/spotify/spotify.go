package spotify

import (
	"context"
	"net/http"

	"github.com/jasperspahl/satpl/pkg/config"
	"golang.org/x/oauth2"
	sptfy "golang.org/x/oauth2/spotify"
)

var oauth2Config *oauth2.Config

func init() {
	oauth2Config = &oauth2.Config{
		ClientID:     config.AppConfig.SpotifyClientId,
		ClientSecret: config.AppConfig.SpotifyClientSecret,
		RedirectURL:  config.AppConfig.SpotifyRedirectUrl,
		Scopes:       []string{"user-read-private", "user-read-email"},
		Endpoint:     sptfy.Endpoint,
	}
}

func AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return oauth2Config.AuthCodeURL(state, opts...)
}

func Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return oauth2Config.Exchange(ctx, code, opts...)
}

type SpotifyClient interface {
	GetCurrentUser() (UserProfile, error)
}

type client struct {
	httpClient *http.Client
}

func Client(ctx context.Context, t *oauth2.Token) SpotifyClient {
	httpClient := oauth2Config.Client(ctx, t)
	return &client{
		httpClient,
	}
}
