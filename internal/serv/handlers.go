package serv

import (
	"encoding/json"
	"net/http"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/info"
	"github.com/sarpt/gamedbv/internal/search"
	plat "github.com/sarpt/gamedbv/pkg/platform"
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
			Platforms: plat.All(),
			Page:      page,
			PageLimit: pageLimit,
		}

		result, err := search.FindGames(appConf, params)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		response := mapToGamesResponse(result)
		out, err := json.Marshal(response)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Write(out)
	}
}

func getLanguagesHandler(appConf config.App) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		result, err := info.Languages(appConf.Database())
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		response := mapToLanguagesResponse(result)
		out, err := json.Marshal(response)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Write(out)
	}
}

func getPlatformsHandler(appConf config.App) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		result, err := info.Platforms(appConf.Database())
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		response := mapToPlatformsResponse(result)
		out, err := json.Marshal(response)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Write(out)
	}
}

func getRegionsHandler(appConf config.App) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		result, err := info.Regions(appConf.Database())
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		response := mapToRegionsResponse(result)
		out, err := json.Marshal(response)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Write(out)
	}
}
