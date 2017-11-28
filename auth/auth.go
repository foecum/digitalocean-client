package auth

import (
	"context"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// TokenSource ...
type TokenSource struct {
	AccessToken string
}

// Token to satisfy the oauth2 Token interface. Returns a Token
func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// NewClient for authentication
func NewClient(accessToken string) *godo.Client {
	tokenSource := &TokenSource{
		AccessToken: accessToken,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	return godo.NewClient(oauthClient)
}
