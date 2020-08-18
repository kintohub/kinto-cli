package api

import (
	"context"
	"github.com/kintohub/kinto-cli/internal/types"
)

func (a *Api) Register(email, password string) (string, error) {
	resp, err := a.authClient.Register(context.Background(), &types.RegisterRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return resp.Token, nil
}

func (a *Api) Login(email, password string) (string, error) {
	resp, err := a.authClient.Login(context.Background(), &types.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return "", err
	}

	return resp.Token, nil
}
