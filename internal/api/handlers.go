package api

import (
	"encoding/json"
	"net/http"

	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/internal/info"
	"github.com/sarpt/gamedbv/pkg/platform"
)

func getGamesHandler(cfg Config) http.HandlerFunc {
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

		params := games.SearchParameters{
			Text:      getTextQuery(req),
			Regions:   getRegionsQuery(req),
			Platforms: platform.All(),
			Page:      page,
			PageLimit: pageLimit,
		}

		result, err := games.Search(cfg.GamesConfig, params)
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

func getLanguagesHandler(cfg Config) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		result, err := info.Languages(cfg.GamesConfig.Database)
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

func getPlatformsHandler(cfg Config) http.HandlerFunc {
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

		result, err := info.Platforms(cfg.GamesConfig.Database, params)
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

func getRegionsHandler(cfg Config) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		result, err := info.Regions(cfg.GamesConfig.Database)
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
