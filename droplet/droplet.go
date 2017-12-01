package droplet

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/foecum/digitalocean-client/client"
)

// Droplet ...
type Droplet struct {
	Client *godo.Client
}

// GetRegions ...
func (d *Droplet) GetRegions() ([]godo.Region, error) {
	ctx := context.TODO()
	regions, _, err := d.Client.Regions.List(ctx, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return regions, nil
}

// GetDroplets ...
func (d *Droplet) GetDroplets() ([]godo.Droplet, error) {
	ctx := context.TODO()

	droplets, _, err := d.Client.Droplets.List(ctx, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return droplets, nil
}
