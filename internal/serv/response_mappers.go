package serv

import (
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/db/queries"
)

func mapToGamesResponse(result queries.GamesResult) gamesResponse {
	games := []gameResponse{}

	for _, item := range result.Games {
		var descriptions []descriptionResponse
		for _, desc := range item.Descriptions {
			descriptions = append(descriptions, descriptionResponse{
				Language: languageResponse{
					Code: desc.Language.Code,
					Name: desc.Language.Name,
				},
				Title:    desc.Title,
				Synopsis: desc.Synopsis,
			})
		}

		games = append(games, gameResponse{
			UID:          item.UID,
			SerialNumber: item.SerialNo,
			Region: regionResponse{
				Code: item.Region.Code,
			},
			Platform: platformResponse{
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
	languages := []languageResponse{}
	for _, result := range results {
		languages = append(languages, languageResponse{
			Code: result.Code,
			Name: result.Name,
		})
	}

	return languagesResponse{
		Languages: languages,
	}
}

func mapToPlatformsResponse(results []models.Platform) platformsResponse {
	platforms := []platformResponse{}
	for _, result := range results {
		platforms = append(platforms, platformResponse{
			UID:  result.UID,
			Name: result.Name,
		})
	}

	return platformsResponse{
		Platforms: platforms,
	}
}

func mapToRegionsResponse(results []models.Region) regionsResponse {
	regions := []regionResponse{}
	for _, result := range results {
		regions = append(regions, regionResponse{
			Code: result.Code,
		})
	}

	return regionsResponse{
		Regions: regions,
	}
}
