package api

import (
	"context"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (a *Api) GetBlocks(envId string) ([]*types.Block, error) {
	env, err := a.GetClusterEnvironment(envId)

	if err != nil {
		utils.TerminateWithError(err)
	}

	blocksResp, err := a.getKubeCoreService(env.ClusterId, envId).GetBlocks(
		context.Background(), &types.BlockQueryRequest{
			EnvId: envId,
		},
	)

	if err != nil {
		return nil, err
	}

	return blocksResp.Items, nil
}
