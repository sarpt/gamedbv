package api

import (
	"github.com/sarpt/gamedbv/pkg/db/models"
	"github.com/sarpt/gamedbv/pkg/db/queries"
)

func mapToGamesResponse(result queries.GamesResult) gamesResponse {
	games := []gameResponse{}

	for _, game := range result.Games {
		var descriptions []descriptionResponse
		for _, desc := range game.Descriptions {
			descriptions = append(descriptions, descriptionResponse{
				Language: languageResponse{
					Code: desc.Language.Code,
					Name: desc.Language.Name,
				},
				Title:    desc.Title,
				Synopsis: desc.Synopsis,
			})
		}

		var regionResponses []regionResponse
		for _, gameRegion := range game.GameRegions {
			regionResponses = append(regionResponses, regionResponse{
				Code: gameRegion.Region.Code,
			})
		}

		games = append(games, gameResponse{
			UID:          game.UID,
			SerialNumber: game.SerialNo,
			Regions:      regionResponses,
			Platform: platformResponse{
				UID: game.Platform.UID,
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
			UID:     result.UID,
			Name:    result.Name,
			Indexed: result.Indexed,
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
