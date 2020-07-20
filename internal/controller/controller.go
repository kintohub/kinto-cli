package controller

import (
	"github.com/kintohub/kinto-cli-go/internal/config"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
	utilsGrpc "github.com/kintohub/utils-go/server/grpc"
)

type ControllerInterface interface {
	Init()
	Register()
	Login()
	Version()

}

type Controller struct {
	authClient     enterpriseTypes.AuthServiceClient
	clustersClient enterpriseTypes.ClusterServiceClient
}

func InitControllerOrDie() ControllerInterface {
	return &Controller{
		authClient:     enterpriseTypes.NewAuthServiceClient(utilsGrpc.CreateConnectionOrDie(config.KintoMasterHost, true)),
		clustersClient: enterpriseTypes.NewClusterServiceClient(utilsGrpc.CreateConnectionOrDie(config.KintoMasterHost, true)),
	}
}
