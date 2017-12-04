package image

import (
	"context"

	"github.com/digitalocean/godo"
	"github.com/foecum/digitalocean-client/client"
)

// Image ...
type Image struct {
	Client *godo.Client
}

// GetImageList ...
func (i *Image) GetImageList() ([]godo.Image, error) {
	ctx := context.TODO()
	images, _, err := i.Client.Images.List(ctx, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return images, nil
}

// GetImageDistributionList ...
func (i *Image) GetImageDistributionList() ([]godo.Image, error) {
	ctx := context.TODO()
	images, _, err := i.Client.Images.ListDistribution(ctx, client.GetOptions())

	if err != nil {
		return nil, err
	}

	return images, nil
}
