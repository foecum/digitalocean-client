package droplet

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/foecum/digitalocean-client/auth"
)

// Droplet ...
type Droplet struct {
	doClient *godo.Client
}

// RegisterClient ...
func (d *Droplet) RegisterClient(accessToken string) {
	client := &auth.Client{}
	d.doClient = client.GetClient(accessToken)
}

// GetRegions ...
func (d *Droplet) GetRegions() ([]godo.Region, error) {
	ctx := context.TODO()

	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	regions, _, err := d.doClient.Regions.List(ctx, opt)

	if err != nil {
		return nil, err
	}

	return regions, nil
}
