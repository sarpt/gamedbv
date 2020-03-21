package serv

import (
	"encoding/json"
	"net/http"
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
	r := mux.NewRouter()

	r.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		params := search.Settings{
			Text:      getTextQueryFromRequest(r),
			Regions:   []string{},
			Platforms: platform.GetAllPlatforms(),
		}

		result, err := search.FindGames(appConf, params)
		if err != nil {
			panic(err)
		}

		out, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	})

	return r
}

func getTextQueryFromRequest(r *http.Request) string {
	return r.URL.Query().Get("q")
}
