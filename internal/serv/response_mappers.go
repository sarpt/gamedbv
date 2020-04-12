package serv

import (
	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/db/models"
)

func mapToGamesResponse(result db.GamesResult) gamesResponse {
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
				Code: item.Platform.Code,
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
			Code: result.Code,
		})
	}

	return platformsResponse{
		Platforms: platforms,
	}
}
