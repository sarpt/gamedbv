package serv

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarpt/gamedbv/internal/config"
)

type handlerCreator func(config.App) http.HandlerFunc

var handlersCreators = map[string]handlerCreator{
	"/games": getGamesHandler,
}

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
	router.Use(corsMiddleware)

	for path, handler := range handlersCreators {
		router.HandleFunc(path, handler(appConf))
	}

	return router
}
