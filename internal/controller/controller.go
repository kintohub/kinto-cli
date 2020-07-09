package controller

import (
	"github.com/kintohub/kinto-cli-go/internal/consts"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
	utilsGrpc "github.com/kintohub/utils-go/server/grpc"
)

type ControllerInterface interface {
	Init()
	Logout()
	Version()
}

type Controller struct {
	authClient     enterpriseTypes.AuthServiceClient
	clustersClient enterpriseTypes.ClusterServiceClient
}

func InitControllerOrDie() ControllerInterface {
	return &Controller{
		authClient:     enterpriseTypes.NewAuthServiceClient(utilsGrpc.CreateConnectionOrDie(consts.KintoMasterHost, true)),
		clustersClient: enterpriseTypes.NewClusterServiceClient(utilsGrpc.CreateConnectionOrDie(consts.KintoMasterHost, true)),
	}
}
