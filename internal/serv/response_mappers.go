package serv

import (
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/db/queries"
)

func mapToGamesResponse(result queries.GamesResult) gamesResponse {
	games := []game{}

	for _, item := range result.Games {
		var descriptions []description
		for _, desc := range item.Descriptions {
			descriptions = append(descriptions, description{
				Language: language{
					Code: desc.Language.Code,
					Name: desc.Language.Name,
				},
				Title:    desc.Title,
				Synopsis: desc.Synopsis,
			})
		}

		games = append(games, game{
			UID:          item.UID,
			SerialNumber: item.SerialNo,
			Region: region{
				Code: item.Region.Code,
			},
			Platform: platform{
				UID: item.Platform.UID,
			},
			Descriptions: descriptions,
		})
	}

	return gamesResponse{
		Total: result.Total,
		Games: games,
	}
}

func mapToLanguagesResponse(results []models.Language) languagesResponse {
	languages := []language{}
	for _, result := range results {
		languages = append(languages, language{
			Code: result.Code,
			Name: result.Name,
		})
	}

	return languagesResponse{
		Languages: languages,
	}
}

func mapToPlatformsResponse(results []models.Platform) platformsResponse {
	platforms := []platform{}
	for _, result := range results {
		platforms = append(platforms, platform{
			UID: result.UID,
		})
	}

	return platformsResponse{
		Platforms: platforms,
	}
}

func mapToRegionsResponse(results []models.Region) regionsResponse {
	regions := []region{}
	for _, result := range results {
		regions = append(regions, region{
			Code: result.Code,
		})
	}

	return regionsResponse{
		Regions: regions,
	}
}
