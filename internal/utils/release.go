package utils

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/kintohub/kinto-kube-core/pkg/types"
	"github.com/rs/zerolog/log"
)

func GetLatestSuccessfulRelease(releases map[string]*types.Release) *types.Release {
	if releases == nil || len(releases) == 0 {
		return nil
	}

	var latestRelease *types.Release
	for _, release := range releases {
		if release.Status.State == types.Status_SUCCESS {
			if latestRelease == nil {
				latestRelease = release
				continue
			}

			latestCreatedAt, err := ptypes.Timestamp(latestRelease.CreatedAt)

			if err != nil {
				log.Error().Err(err).Msgf(
					"cannot parse timestamp %v to time for release %v",
					latestRelease.CreatedAt, latestRelease)
				continue
			}

			releaseCreateAt, err := ptypes.Timestamp(release.CreatedAt)

			if err != nil {
				log.Error().Err(err).Msgf(
					"cannot parse timestamp %v to time for release %v",
					release.CreatedAt, release)
				continue
			}

			if releaseCreateAt.After(latestCreatedAt) {
				latestRelease = release
			}
		}
	}

	return latestRelease
}
