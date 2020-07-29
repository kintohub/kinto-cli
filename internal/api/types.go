package api

import (
	"context"
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
	"google.golang.org/grpc/metadata"
)

// Implements grpc PerRPCCredentials
type accessTokenManager struct {
	clusterId      string
	envId          string
	authToken      string
	clustersClient enterpriseTypes.ClusterServiceClient
}

func (a *accessTokenManager) GetRequestMetadata(ctx context.Context, args ...string) (map[string]string, error) {
	// TODO: store + check if expired logic
	bearer := "Bearer " + config.GetAuthToken()
	md := metadata.Pairs("Authorization", bearer)
	tokenCtx := metadata.NewOutgoingContext(ctx, md)

	token, err := a.clustersClient.CreateAccessToken(tokenCtx, &enterpriseTypes.CreateAccessTokenRequest{
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
