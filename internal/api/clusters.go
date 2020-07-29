package api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/kintohub/kinto-cli-go/internal/config"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
	"google.golang.org/grpc/metadata"
)

func (a *Api) GetClusterEnvironment(envId string) (*enterpriseTypes.ClusterEnvironment, error) {
	env := config.GetClusterEnvironment(envId)

	if env == nil {
		envs, err := a.GetClusterEnvironments()

		if err != nil {
			return nil, err
		}

		for _, env := range envs {
			if env.Id == envId {
				return env, nil
			}
		}

		return nil, NotFoundError
	}

	return env, nil
}

func (a *Api) GetClusterEnvironments() ([]*enterpriseTypes.ClusterEnvironment, error) {
	bearer := "Bearer " + config.GetAuthToken()
	md := metadata.Pairs("Authorization", bearer)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := a.clusterClient.GetClusterEnvironments(ctx, &empty.Empty{})

	if err != nil {
		return nil, err
	}

	return resp.Envs, err
}

func (a *Api) GetClusters() ([]*enterpriseTypes.PublicClusterInfo, error) {
	resp, err := a.clusterClient.GetClusters(context.Background(), &empty.Empty{})

	if err != nil {
		return nil, err
	}

	config.SetPublicClusters(resp.Clusters)

	return resp.Clusters, err
}

func (a *Api) GetPublicClusterInfo(clusterId string) (*enterpriseTypes.PublicClusterInfo, error) {
	publicClusterInfo := config.GetPublicClusterInfo(clusterId)

	if publicClusterInfo == nil {
		clusters, err := a.GetClusters()

		if err != nil {
			return nil, err
		}

		for _, cluster := range clusters {
			if cluster.Id == clusterId {
				return cluster, nil
			}
		}

		return nil, NotFoundError
	}

	return publicClusterInfo, nil
}
