package api

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarpt/gamedbv/internal/games"
)

// Config instructs how API should behave and how it should access indexes and database
type Config struct {
	IPAddress    string
	Port         string
	NetInterface string
	Debug        bool
	GamesConfig  games.Config
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
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
func Serve(cfg Config, w io.Writer) error {
	router := initRouter(cfg)

	address := addressForServe(cfg)
	fmt.Fprintf(w, "API server address: %s\n", address)

	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
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

func addressForServe(cfg Config) string {
	var ip string = cfg.IPAddress
	port := cfg.Port

	if cfg.NetInterface != "" {
		foundIP, err := findIP(cfg.NetInterface)
		if err == nil {
			ip = foundIP
		}
	}

	return strings.Join([]string{ip, port}, ":")
}

func findIP(interfaceName string) (string, error) {
	netIf, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", err
	}

	ifAddrs, err := netIf.Addrs()
	if err != nil {
		return "", err
	}

	for _, ifAddr := range ifAddrs {
		var ip *net.IP
		switch addr := ifAddr.(type) {
		case *net.IPNet:
			ip = &addr.IP
		case *net.IPAddr:
			ip = &addr.IP
		default:
			return "", err
		}

		if ip == nil || ip.IsLoopback() || ip.To4() == nil {
			continue
		}

		return ip.String(), nil
	}

	return "", fmt.Errorf("no suitable ipv4 address found for %s interface", interfaceName)
}
