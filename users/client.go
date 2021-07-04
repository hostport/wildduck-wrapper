package users

import (
	"errors"
	"fmt"
	"github.com/hostport/wildduck-wrapper"
	"net/http"
)

type Client struct {
	Backend wildduck.Backend
}

func New(user *wildduck.UserParams) error {
	return getClient().New(user)
}

func (c Client) New(user *wildduck.UserParams) error {
	res := &struct {
		Success bool   `json:"success"`
		Id      string `json:"id"`
	}{}
	err := c.Backend.Call(http.MethodPost, "users", user, res)
	if err != nil {
		return err
	}
	if !res.Success {
		return errors.New("could not create user")
	}
	user.Id = res.Id
	return nil
}

func Update(userId string, user *wildduck.UserParams) error {
	return getClient().Update(userId, user)
}

func (c Client) Update(userId string, user *wildduck.UserParams) error {
	res := struct {
		Success bool `json:"success"`
	}{}
	err := c.Backend.Call(http.MethodPut, fmt.Sprintf("users/%s", userId), user, &res)
	if err != nil {
		return err
	}
	if !res.Success {
		return errors.New("could not update user")
	}
	return nil
}

func Delete(userId string) error {
	return getClient().Delete(userId)
}

func (c Client) Delete(userId string) error {
	res := struct {
		Success bool `json:"success"`
	}{}
	err := c.Backend.Call(http.MethodDelete, fmt.Sprintf("users/%s", userId), nil, &res)
	if err != nil {
		return err
	}
	if !res.Success {
		return errors.New("could not delete user")
	}
	return nil
}

func GetAll() (*wildduck.AllUsersResponse, error) {
	return getClient().GetAll()
}

func (c Client) GetAll() (*wildduck.AllUsersResponse, error) {
	res := &wildduck.AllUsersResponse{}
	err := c.Backend.Call(http.MethodGet, "users", nil, res)
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errors.New("could not get users")
	}
	return res, nil
}

func getClient() Client {
	return Client{wildduck.GetBackend()}
}
