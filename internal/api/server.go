package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/internal/server"
)

const (
	loggerPrefox = "api.Server#"
)

// Config instructs how API should behave and how it should access indexes and database
type Config struct {
	Debug          bool
	DlRPCAddress   string
	DlRPCPort      string
	ErrWriter      io.Writer
	GamesConfig    games.Config
	IdxRPCAddress  string
	IdxRPCPort     string
	IPAddress      string
	NetInterface   string
	OutWriter      io.Writer
	Port           string
	ReadTimeout    time.Duration
	RPCDialTimeout time.Duration
	StartServices  bool
	WriteTimeout   time.Duration
}

// Server represents API server instance
type Server struct {
	cfg               Config
	dlService         *DlService
	errLog            *log.Logger
	idxService        *IdxService
	operationHandlers map[operation]operationHandler
	outLog            *log.Logger
	routeHandlers     map[string]http.HandlerFunc
}

// NewServer returns new API server instance.
func NewServer(cfg Config) *Server {
	server := Server{
		cfg:        cfg,
		dlService:  NewDlService(cfg),
		errLog:     log.New(cfg.ErrWriter, loggerPrefox, log.LstdFlags),
		idxService: NewIdxService(cfg),
		outLog:     log.New(cfg.OutWriter, loggerPrefox, log.LstdFlags),
	}
	server.routeHandlers = server.getRouteHandlers()
	server.operationHandlers = server.getOperationHandlers()

	return &server
}

// Serve starts GameDBV API server.
func (s *Server) Serve(out io.Writer) error {
	err := s.dialGRPCServices()
	if err != nil {
		return fmt.Errorf("could not dial GRPC services: %w", err)
	}
	defer s.closeGRPCConnections()

	router := s.initRouter()

	address := s.addressForServe()
	s.outLog.Printf("API server address: %s\n", address)

	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: s.cfg.WriteTimeout,
		ReadTimeout:  s.cfg.ReadTimeout,
	}

	return srv.ListenAndServe()
}

func (s Server) initRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(corsMiddleware)
	router.Use(jsonAPIMiddleware)

	for path, handler := range s.routeHandlers {
		router.HandleFunc(path, handler)
	}

	return router
}

func (s Server) addressForServe() string {
	var ip string = s.cfg.IPAddress
	port := s.cfg.Port

	if s.cfg.NetInterface != "" {
		foundIP, err := server.IPByInterfaceName(s.cfg.NetInterface)
		if err == nil {
			ip = foundIP
		}
	}

	return strings.Join([]string{ip, port}, ":")
}
