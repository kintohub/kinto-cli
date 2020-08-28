package api

import (
	"crypto/x509"
	"errors"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/kintohub/kinto-cli/internal/utils"
	utilsGrpc "github.com/kintohub/utils-go/server/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"time"
)

var (
	NotFoundError = errors.New("resource requested was not found")
)

type ApiInterface interface {
	GetClusterEnvironments() ([]*types.ClusterEnvironment, error)
	GetClusters() ([]*types.PublicClusterInfo, error)
	Login(email, password string) (string, error)
	GetBlocks(envId string) ([]*types.Block, error)
	StartTeleport(blocksToForward []RemoteConfig, envId string, clusterId string)
}

// Due to the nature of APIs,
type Api struct {
	masterHost             string
	authClient             types.AuthServiceClient
	clusterClient          types.ClusterServiceClient
	kubeCoreServiceClients map[string]types.KintoKubeCoreServiceClient
}

func NewApiOrDie(masterHost string) ApiInterface {
	return &Api{
		masterHost:             masterHost,
		authClient:             types.NewAuthServiceClient(utilsGrpc.CreateConnectionOrDie(masterHost, true)),
		clusterClient:          types.NewClusterServiceClient(utilsGrpc.CreateConnectionOrDie(masterHost, true)),
		kubeCoreServiceClients: map[string]types.KintoKubeCoreServiceClient{},
	}
}

func (a *Api) getKubeCoreService(clusterId, envId string) types.KintoKubeCoreServiceClient {
	publicCluster, err := a.GetPublicClusterInfo(clusterId)

	if err != nil {
		utils.TerminateWithError(err)
	}

	if service, ok := a.kubeCoreServiceClients[publicCluster.Id]; ok {
		return service
	} else {
		client := createKintoKubeCoreClientOrDie(
			a.clusterClient,
			publicCluster,
			envId,
		)

		a.kubeCoreServiceClients[publicCluster.Id+envId] = client
		return client
	}
}

// TODO: Refactor kinto go commons to accept 3rd party dial options optionally
func createKintoKubeCoreClientOrDie(
	clustersClient types.ClusterServiceClient,
	cluster *types.PublicClusterInfo,
	envId string) types.KintoKubeCoreServiceClient {
	// https://grpc.io/docs/guides/auth/#authenticate-with-google
	pool, _ := x509.SystemCertPool()
	creds := credentials.NewClientTLSFromCert(pool, "")
	dialOption := grpc.WithTransportCredentials(creds)

	conn, err := grpc.Dial(
		cluster.HostName+":443",
		grpc.WithPerRPCCredentials(&accessTokenManager{
			envId:          envId,
			clusterId:      cluster.Id,
			authToken:      config.GetAuthToken(),
			clustersClient: clustersClient,
		}),
		dialOption,
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
			Timeout:             time.Second * 5,  // wait 1 second for ping ack before considering the connection dead
			PermitWithoutStream: true,             // send pings even without active streams
		}),
	)

	if err != nil {
		utils.TerminateWithError(err)
	}

	return types.NewKintoKubeCoreServiceClient(conn)
}
