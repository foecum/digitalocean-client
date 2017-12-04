package droplet

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

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

// GetAllDroplets ...
func (d *Droplet) GetAllDroplets() ([]godo.Droplet, error) {
	ctx := context.TODO()
	droplets, _, err := d.Client.Droplets.List(ctx, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return droplets, nil
}

// GetDropletByID ...
func (d *Droplet) GetDropletByID(dropletID int) (*godo.Droplet, error) {
	ctx := context.TODO()

	droplet, _, err := d.Client.Droplets.Get(ctx, dropletID)

	if err != nil {
		return nil, err
	}

	return droplet, nil
}

// GetDropletsByTag ...
func (d *Droplet) GetDropletsByTag(tag string) ([]godo.Droplet, error) {
	ctx := context.TODO()

	droplets, _, err := d.Client.Droplets.ListByTag(ctx, tag, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return droplets, nil
}

// GetDropletAvailableKernels ...
func (d *Droplet) GetDropletAvailableKernels(dropletID int) ([]godo.Kernel, error) {
	ctx := context.TODO()

	kernels, _, err := d.Client.Droplets.Kernels(ctx, dropletID, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return kernels, nil
}

// GetDropletSnapshots ...
func (d *Droplet) GetDropletSnapshots(dropletID int) ([]godo.Image, error) {
	ctx := context.TODO()

	snapshots, _, err := d.Client.Droplets.Snapshots(ctx, dropletID, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return snapshots, nil
}

// GetDropletBackups ...
func (d *Droplet) GetDropletBackups(dropletID int) ([]godo.Image, error) {
	ctx := context.TODO()

	backups, _, err := d.Client.Droplets.Backups(ctx, dropletID, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return backups, nil
}

// GetDropletActions ...
func (d *Droplet) GetDropletActions(dropletID int) ([]godo.Action, error) {
	ctx := context.TODO()

	actions, _, err := d.Client.Droplets.Actions(ctx, dropletID, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return actions, nil
}

// DeleteDroplet ...
func (d *Droplet) DeleteDroplet(dropletID int) error {
	ctx := context.TODO()

	_, err := d.Client.Droplets.Delete(ctx, dropletID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteDropletByTag ...
func (d *Droplet) DeleteDropletByTag(tag string) error {
	ctx := context.TODO()

	_, err := d.Client.Droplets.DeleteByTag(ctx, tag)

	if err != nil {
		return err
	}

	return nil
}

// GetDropletNeighbors ...
func (d *Droplet) GetDropletNeighbors(dropletID int) ([]byte, error) {
	ctx := context.TODO()

	endpoint := fmt.Sprintf("/v2/droplets/%d/neighbors", dropletID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := d.Client.Do(ctx, req, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetAllDropletNeighbors ...
func (d *Droplet) GetAllDropletNeighbors() ([]byte, error) {
	ctx := context.TODO()

	endpoint := "/v2/reports/droplet_neighbors"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := d.Client.Do(ctx, req, nil)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return b, nil
}
