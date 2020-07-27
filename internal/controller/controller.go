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
	Environment()
	Version()
}

type Controller struct {
	authClient     enterpriseTypes.AuthServiceClient
	clustersClient enterpriseTypes.ClusterServiceClient
}

func InitControllerOrDie() ControllerInterface {

	// Can't find a valid way to make MasterKintoHost as a flag parameter since Controller
	// is initialized before CLI

	// TODO : Make MasterKintoHost as flag parameter
	return &Controller{
		authClient: enterpriseTypes.
			NewAuthServiceClient(utilsGrpc.CreateConnectionOrDie(config.KintoMasterHost,
				true)),
		clustersClient: enterpriseTypes.
			NewClusterServiceClient(utilsGrpc.CreateConnectionOrDie(config.KintoMasterHost,
				true)),
	}
}
