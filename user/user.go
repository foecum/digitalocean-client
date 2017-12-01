package user

import (
	"context"

	"github.com/digitalocean/godo"
)

// User ...
type User struct {
	Client *godo.Client
}

// GetUserInfo ...
func (u *User) GetUserInfo() (godo.Account, error) {
	ctx := context.TODO()
	account, _, err := u.Client.Account.Get(ctx)

	if err != nil {
		return godo.Account{}, err
	}

	return *account, nil
}
