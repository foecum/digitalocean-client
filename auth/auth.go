package auth

import (
	"context"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// Client ...
type Client struct {
	cl *godo.Client
}

// TokenSource ...
type tokenSource struct {
	AccessToken string
}

// Token to satisfy the oauth2 Token interface. Returns a Token
func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

// GetClient for authentication
func (c *Client) GetClient(accessToken string) *godo.Client {
	if c.cl == nil {
		tokenSource := &tokenSource{
			AccessToken: accessToken,
		}

		oauthClient := oauth2.NewClient(context.Background(), tokenSource)
		c.cl = godo.NewClient(oauthClient)
	}
	return c.cl
}
