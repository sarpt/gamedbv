package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/sarpt/gamedbv/internal/cmds"
	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/internal/status"
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

		gamesCfg := cmds.GamesCfg{}
		gamesArgs := cmds.GamesArguments{
			Text:      getTextQuery(req),
			Regions:   getRegionsQuery(req),
			Platforms: getPlatformsQuery(req),
			Page:      page,
			PageLimit: pageLimit,
		}

		gamesCmd := cmds.NewGames(gamesCfg, gamesArgs)

		result, err := gamesCmd.Execute()
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
		result, err := status.Languages(cfg.GamesConfig.Database)
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

		params := status.PlatformsParameters{
			Indexed: filterIndexed,
			UID:     uid,
		}

		result, err := status.Platforms(cfg.GamesConfig.Database, params)
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
		result, err := status.Regions(cfg.GamesConfig.Database)
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

		for {
			cmdMsg := clientOpertionMessage{}
			err := conn.ReadJSON(&cmdMsg)
			if err != nil {
				var closeError *websocket.CloseError
				if errors.As(err, &closeError) {
					fmt.Fprintf(os.Stderr, "Connection was closed: %s\n", err)

					break
				}

				status := operationStatus{
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

			go func() {
				sw := newProgressWriter(conn)
				err = handleOperationMessage(cmdMsg, sw)
				if err != nil {
					status := operationStatus{
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
					status := operationStatus{
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
			}()
		}
	}
}
