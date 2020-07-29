package api

import (
	"context"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	kkcTypes "github.com/kintohub/kinto-kube-core/pkg/types"
)

func (a *Api) GetBlocks(envId string) ([]*kkcTypes.Block, error) {
	env, err := a.GetClusterEnvironment(envId)

	if err != nil {
		utils.TerminateWithError(err)
	}

	blocksResp, err := a.getKubeCoreService(env.ClusterId, envId).GetBlocks(
		context.Background(), &kkcTypes.BlockQueryRequest{
			EnvId: envId,
		},
	)

	if err != nil {
		return nil, err
	}

	return blocksResp.Items, nil
}
