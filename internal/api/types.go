package api

import (
	"context"
	"fmt"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/kintohub/kinto-cli/internal/utils"
	"google.golang.org/grpc/metadata"
)

// Implements grpc PerRPCCredentials
type accessTokenManager struct {
	clusterId      string
	envId          string
	authToken      string
	clustersClient types.ClusterServiceClient
}

type RemoteConfig struct {
	FromHost string
	FromPort int
	ToHost   string
	ToPort   int
}

func (a *accessTokenManager) GetRequestMetadata(ctx context.Context, args ...string) (map[string]string, error) {
	// TODO: store + check if expired logic
	bearer := "Bearer " + config.GetAuthToken()
	md := metadata.Pairs("Authorization", bearer)
	tokenCtx := metadata.NewOutgoingContext(ctx, md)

	token, err := a.clustersClient.CreateAccessToken(tokenCtx, &types.CreateAccessTokenRequest{
		ClusterId: a.clusterId,
		EnvId:     a.envId,
	})

	if err != nil {
		utils.TerminateWithError(err)
	}

	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
	}, nil
}

func (a *accessTokenManager) RequireTransportSecurity() bool {
	return true
}
