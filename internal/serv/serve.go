package serv

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/search"
	"github.com/sarpt/gamedbv/pkg/platform"
)

// Serve starts GameDBV server
func Serve(appConf config.App) error {
	router := initRouter(appConf)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func initRouter(appConf config.App) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/games", func(resp http.ResponseWriter, req *http.Request) {
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
			Regions:   []string{},
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
	})

	return router
}

func getTextQueryFromRequest(r *http.Request) string {
	return r.URL.Query().Get("q")
}

func getCurrentPageFromRequest(r *http.Request) (int, error) {
	page := r.URL.Query().Get("_page")
	if page == "" {
		return 0, nil
	}

	return strconv.Atoi(page)
}

func getPageLimitFromRequest(r *http.Request) (int, error) {
	limit := r.URL.Query().Get("_limit")
	if limit == "" {
		return -1, nil
	}

	return strconv.Atoi(limit)
}

func getPlatformsFromRequest(r *http.Request) []string {
	query := r.URL.Query()
	if platforms, ok := query["platform"]; ok {
		return platforms
	}

	return []string{}
}

func getRegionsFromRequest(r *http.Request) []string {
	query := r.URL.Query()
	if regions, ok := query["region"]; ok {
		return regions
	}

	return []string{}
}
