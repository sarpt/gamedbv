package serv

import "github.com/sarpt/gamedbv/pkg/db"

func mapToGamesResponse(result db.GamesResult) gamesResponse {
	var games []game
	for _, item := range result.Games {
		var descriptions []description
		for _, desc := range item.Descriptions {
			descriptions = append(descriptions, description{
				Language: desc.Language.Code,
				Title:    desc.Title,
				Synopsis: desc.Synopsis,
			})
		}

		games = append(games, game{
			UID:          item.UID,
			SerialNumber: item.SerialNo,
			Region:       item.Region,
			Platform:     item.Platform.Name,
			Descriptions: descriptions,
		})
	}

	return gamesResponse{
		Total: result.Total,
		Games: games,
	}
}
