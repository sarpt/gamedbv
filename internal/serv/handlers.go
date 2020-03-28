package serv

import (
	"encoding/json"
	"net/http"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/search"
	"github.com/sarpt/gamedbv/pkg/platform"
)

func getGamesHandler(appConf config.App) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		page, err := getCurrentPageFromRequest(req)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		pageLimit, err := getPageLimitFromRequest(req)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		params := search.Settings{
			Text:      getTextQueryFromRequest(req),
			Regions:   getRegionsFromRequest(req),
			Platforms: platform.GetAllPlatforms(),
			Page:      page,
			PageLimit: pageLimit,
		}

		result, err := search.FindGames(appConf, params)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		out, err := json.Marshal(result)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Header().Set("Content-Type", "application/json")
		resp.Write(out)
	}
}
