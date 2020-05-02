package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarpt/gamedbv/internal/games"
)

// Config instructs how API should behave and how it should access indexes and database
type Config struct {
	GamesConfig games.Config
}

type handlerCreator func(Config) http.HandlerFunc

var handlersCreators = map[string]handlerCreator{
	"/games":          getGamesHandler,
	"/info/languages": getLanguagesHandler,
	"/info/platforms": getPlatformsHandler,
	"/info/regions":   getRegionsHandler,
	"/updates":        getUpdatesHandler,
}

// Serve starts GameDBV API server
func Serve(cfg Config) error {
	router := initRouter(cfg)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func initRouter(cfg Config) *mux.Router {
	router := mux.NewRouter()
	router.Use(corsMiddleware)
	router.Use(jsonAPIMiddleware)

	for path, handler := range handlersCreators {
		router.HandleFunc(path, handler(cfg))
	}

	return router
}
