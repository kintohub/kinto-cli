package api

import (
	"crypto/x509"
	"errors"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
	kkcTypes "github.com/kintohub/kinto-kube-core/pkg/types"
	utilsGrpc "github.com/kintohub/utils-go/server/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	NotFoundError = errors.New("resource requested was not found")
)

type ApiInterface interface {
	GetClusterEnvironments() ([]*enterpriseTypes.ClusterEnvironment, error)
	GetClusters() ([]*enterpriseTypes.PublicClusterInfo, error)
	Register(email, password string) (string, error)
	Login(email, password string) (string, error)
}

// Due to the nature of APIs,
type Api struct {
	masterHost             string
	authClient             enterpriseTypes.AuthServiceClient
	clusterClient          enterpriseTypes.ClusterServiceClient
	kubeCoreServiceClients map[string]kkcTypes.KintoKubeCoreServiceClient
}

func NewApiOrDie(masterHost string) ApiInterface {
	return &Api{
		masterHost: masterHost,
		authClient: enterpriseTypes.
			NewAuthServiceClient(utilsGrpc.CreateConnectionOrDie(masterHost, true)),
		clusterClient: enterpriseTypes.
			NewClusterServiceClient(utilsGrpc.CreateConnectionOrDie(masterHost, true)),
	}
}

func (a *Api) getKubeCoreService(clusterId string) kkcTypes.KintoKubeCoreServiceClient {
	publicCluster, err := a.GetPublicClusterInfo(clusterId)

	if err != nil {
		utils.TerminateWithError(err)
	}

	if service, ok := a.kubeCoreServiceClients[publicCluster.Id]; ok {
		return service
	} else {
		client := createKintoKubeCoreClientOrDie(publicCluster.HostName)
		a.kubeCoreServiceClients[publicCluster.Id] = client
		return client
	}
}

// TODO: Refactor kinto go commons to accept 3rd party dial options optionally
func createKintoKubeCoreClientOrDie(clusterHost string) kkcTypes.KintoKubeCoreServiceClient {
	dialOption := grpc.WithInsecure()

	// https://grpc.io/docs/guides/auth/#authenticate-with-google
	pool, _ := x509.SystemCertPool()
	creds := credentials.NewClientTLSFromCert(pool, "")
	dialOption = grpc.WithTransportCredentials(creds)

	conn, err := grpc.Dial(
		clusterHost,
		dialOption,
		grpc.WithPerRPCCredentials(&accessTokenManager{
			authToken: config.GetAuthToken(),
		}),
	)

	if err != nil {
		utils.TerminateWithError(err)
	}

	return kkcTypes.NewKintoKubeCoreServiceClient(conn)
}
