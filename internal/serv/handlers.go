package serv

import (
	"encoding/json"
	"net/http"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/info"
	"github.com/sarpt/gamedbv/internal/search"
	"github.com/sarpt/gamedbv/pkg/platform"
)

func getGamesHandler(appConf config.App) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		page, err := getCurrentPageQuery(req)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		pageLimit, err := getPageLimitQuery(req)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		params := search.Parameters{
			Text:      getTextQuery(req),
			Regions:   getRegionsQuery(req),
			Platforms: platform.All(),
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
		filterIndexed, err := getIndexedQuery(req)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		uid := getUIDQuery(req)

		params := info.PlatformsParameters{
			Indexed: filterIndexed,
			UID:     uid,
		}

		result, err := info.Platforms(appConf.Database(), params)
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
