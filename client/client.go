package client

import (
	"github.com/digitalocean/godo"
	"github.com/foecum/digitalocean-client/auth"
)

// RegisterClient ...
func RegisterClient(accessToken string) *godo.Client {
	client := &auth.Client{}
	return client.GetClient(accessToken)
}

// GetOptions ...
func GetOptions() *godo.ListOptions {
	return &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}
}
