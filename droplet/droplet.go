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

func (d *Droplet) getOptions() *godo.ListOptions {
	return &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}
}

// GetRegions ...
func (d *Droplet) GetRegions() ([]godo.Region, error) {
	ctx := context.TODO()
	regions, _, err := d.doClient.Regions.List(ctx, d.getOptions())

	if err != nil {
		return nil, err
	}

	return regions, nil
}

// GetDroplets ...
func (d *Droplet) GetDroplets() ([]godo.Droplet, error) {
	ctx := context.TODO()

	droplets, _, err := d.doClient.Droplets.List(ctx, d.getOptions())

	if err != nil {
		return nil, err
	}

	return droplets, nil
}
