package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/internal/info"
	"github.com/sarpt/gamedbv/internal/progress"
)

func getGamesHandler(cfg Config) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		page, err := getCurrentPageQuery(req)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		pageLimit, err := getPageLimitQuery(req)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		params := games.SearchParameters{
			Text:      getTextQuery(req),
			Regions:   getRegionsQuery(req),
			Platforms: getPlatformVariants(req),
			Page:      page,
			PageLimit: pageLimit,
		}

		result, err := games.Search(cfg.GamesConfig, params)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		response := mapToGamesResponse(result)
		out, err := json.Marshal(response)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Write(out)
	}
}

func getLanguagesHandler(cfg Config) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		result, err := info.Languages(cfg.GamesConfig.Database)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		response := mapToLanguagesResponse(result)
		out, err := json.Marshal(response)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Write(out)
	}
}

func getPlatformsHandler(cfg Config) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		filterIndexed, err := getIndexedQuery(req)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		uid := getUIDQuery(req)

		params := info.PlatformsParameters{
			Indexed: filterIndexed,
			UID:     uid,
		}

		result, err := info.Platforms(cfg.GamesConfig.Database, params)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		response := mapToPlatformsResponse(result)
		out, err := json.Marshal(response)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Write(out)
	}
}

func getRegionsHandler(cfg Config) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		result, err := info.Regions(cfg.GamesConfig.Database)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		response := mapToRegionsResponse(result)
		out, err := json.Marshal(response)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Write(out)
	}
}

func getUpdatesHandler(cfg Config) http.HandlerFunc {
	upgrader := websocket.Upgrader{}
	if cfg.Debug {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}

	return func(res http.ResponseWriter, req *http.Request) {
		conn, err := upgrader.Upgrade(res, req, nil)
		if err != nil {
			return
		}

		var done chan bool

		go func() {
			for {
				cmdMsg := clientCmdMessage{}
				err := conn.ReadJSON(&cmdMsg)
				if err != nil {
					var closeError *websocket.CloseError
					if errors.As(err, &closeError) {
						fmt.Fprintf(os.Stderr, "Connection was closed: %s\n", err)

						break
					}

					status := statusMessage{
						State: errorState,
						Status: progress.Status{
							Message: err.Error(),
						},
					}
					err = conn.WriteJSON(&status)
					if err != nil {
						fmt.Fprintf(os.Stderr, "err: %s\n", err) // other kind of logging?
					}

					continue
				}

				sw := newProgressWriter(conn)
				err = handleCmdMessage(cmdMsg, sw)
				if err != nil {
					status := statusMessage{
						State: errorState,
						Status: progress.Status{
							Message: err.Error(),
						},
					}
					err = conn.WriteJSON(&status)
					if err != nil {
						fmt.Fprintf(os.Stderr, "err: %s\n", err)
					}
				} else {
					status := statusMessage{
						State: doneState,
						Status: progress.Status{
							Message: "Command finished",
						},
					}
					err = conn.WriteJSON(&status)
					if err != nil {
						fmt.Fprintf(os.Stderr, "err: %s\n", err)
					}
				}
			}

			done <- true
		}()

		<-done
	}
}
