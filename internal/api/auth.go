package api

import (
	"context"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
)

func (a *Api) Register(email, password string) (string, error) {
	resp, err := a.authClient.Register(context.Background(), &enterpriseTypes.RegisterRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (a *Api) Login(email, password string) (string, error) {
	resp, err := a.authClient.Login(context.Background(), &enterpriseTypes.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return resp.Token, nil
}
