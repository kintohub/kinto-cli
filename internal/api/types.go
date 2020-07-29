package api

import (
	"context"
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
)

// Implements grpc PerRPCCredentials
type accessTokenManager struct {
	clusterId          string
	envId              string
	authToken          string
	clusterAccessToken string
	clustersClient     enterpriseTypes.ClusterServiceClient
}

func (a *accessTokenManager) GetRequestMetadata(ctx context.Context, args ...string) (map[string]string, error) {
	if a.clusterAccessToken != "" {
		// TODO: Check if it expired
		token, err := a.clustersClient.CreateAccessToken(ctx, &enterpriseTypes.CreateAccessTokenRequest{
			ClusterId: a.clusterId,
			EnvId:     a.envId,
		})

		if err != nil {
			utils.TerminateWithError(err)
		}

		a.clusterAccessToken = token.AccessToken
	}

	return map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", a.clusterAccessToken),
	}, nil
}

func (a *accessTokenManager) RequireTransportSecurity() bool {
	return true
}
