package image

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

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

// GetApplicationImageList ...
func (i *Image) GetApplicationImageList() ([]godo.Image, error) {
	ctx := context.TODO()
	images, _, err := i.Client.Images.ListApplication(ctx, client.GetOptions())
	if err != nil {
		return nil, err
	}

	return images, nil
}

// GetUserImageList ...
func (i *Image) GetUserImageList() ([]godo.Image, error) {
	ctx := context.TODO()
	images, _, err := i.Client.Images.ListUser(ctx, client.GetOptions())
	if err != nil {
		return nil, err
	}

	return images, nil
}

// GetPrivateUserImageList ...
func (i *Image) GetPrivateUserImageList() ([]byte, error) {
	ctx := context.TODO()

	endpoint := fmt.Sprintf("/v2/images?private=true")

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := i.Client.Do(ctx, req, nil)

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

// GetImageActions ...
func (i *Image) GetImageActions(dropletID int) ([]byte, error) {
	ctx := context.TODO()

	endpoint := fmt.Sprintf("/v2/images/%d/actions", dropletID)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := i.Client.Do(ctx, req, nil)

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

// GetExistingImageByID ...
func (i *Image) GetExistingImageByID(imageID int) (*godo.Image, error) {
	ctx := context.TODO()
	image, _, err := i.Client.Images.GetByID(ctx, imageID)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// GetExistingImageBySlug ...
func (i *Image) GetExistingImageBySlug(slug string) (*godo.Image, error) {
	ctx := context.TODO()
	image, _, err := i.Client.Images.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// UpdateImage ...
func (i *Image) UpdateImage(imageID int, imageName string) (*godo.Image, error) {
	ctx := context.TODO()

	updateRequest := &godo.ImageUpdateRequest{
		Name: imageName,
	}

	image, _, err := i.Client.Images.Update(ctx, imageID, updateRequest)
	if err != nil {
		return nil, err
	}

	return image, nil
}

// DeleteImage ...
func (i *Image) DeleteImage(imageID int) (*godo.Image, error) {
	ctx := context.TODO()

	_, err := i.Client.Images.Delete(ctx, imageID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
