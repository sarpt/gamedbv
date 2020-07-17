package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/internal/status"
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

func handlePlatforms(dbCfg db.Config) (progress.Status, error) {
	var indexingFlag *string
	var uidFlag *string
	var resultStatus progress.Status = progress.Status{}

	flagSet := flag.NewFlagSet("platforms", flag.PanicOnError)
	indexingFlag = flagSet.String("indexing", string(status.AllPlatforms), "possible values: all, with, without")
	uidFlag = flagSet.String("uid", "", "when uid is specified, only platform matching the uid will be returned")
	flagSet.Parse(os.Args)

	var indexed status.FilterIndexing

	switch *indexingFlag {
	case string(status.AllPlatforms):
		indexed = status.AllPlatforms
	case string(status.WithIndex):
		indexed = status.WithIndex
	case string(status.WithoutIndex):
		indexed = status.WithoutIndex
	default:
		return resultStatus, fmt.Errorf("%s is not a correct indexing value", *indexingFlag)
	}
	params := status.PlatformsParameters{
		Indexed: indexed,
		UID:     *uidFlag,
	}

	result, err := status.Platforms(dbCfg, params)
	if err != nil {
		return resultStatus, err
	}

	out := preparePlatformsOutput(result)
	resultStatus = prepareResultStatus(result, out)
	return resultStatus, nil
}

func preparePlatformsOutput(platforms []models.Platform) string {
	var out string = ""

	for _, platform := range platforms {
		out = fmt.Sprintf("%s\n[%s] %s", out, platform.UID, platform.Name)
	}

	return fmt.Sprintf("%s\n", out)
}
