package games

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/queries"
)

func gamesDetailsFromDatabase(dbCfg db.Config, params SearchParameters, serialNumbers []string) (queries.GamesResult, error) {
	var gamesResult queries.GamesResult

	database, err := db.OpenDatabase(dbCfg)
	defer database.Close()

	if err != nil {
		return gamesResult, err
	}

	gamesQuery := database.NewGamesQuery()
	if len(serialNumbers) > 0 {
		gamesQuery.FilterUIDs(serialNumbers)
	}

	if len(params.Platforms) > 0 {
		var platformIds []string

		for _, platVariant := range params.Platforms {
			platformIds = append(platformIds, platVariant.ID())
		}

		gamesQuery.FilterPlatforms(platformIds)
	}

	if len(params.Regions) > 0 {
		gamesQuery.FilterRegions(params.Regions)
	}

	gamesQuery.Page(params.Page)
	gamesQuery.Limit(params.PageLimit)

	gamesResult = gamesQuery.Get()
	return gamesResult, err
}
