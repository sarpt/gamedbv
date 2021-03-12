package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/sarpt/gamedbv/internal/games"
	"github.com/sarpt/gamedbv/internal/server"
	pb "github.com/sarpt/gamedbv/pkg/rpc/dl"
	"google.golang.org/grpc"
)

const (
	gamesEndpoint     = "/games"
	languagesEndpoint = "/info/languages"
	platformsEndpoint = "/info/platforms"
	regionsEndpoint   = "/info/regions"
	updatesEndpoint   = "/updates"
)

// Config instructs how API should behave and how it should access indexes and database
type Config struct {
	Debug          bool
	DlRPCAddress   string
	DlRPCPort      string
	GamesConfig    games.Config
	IPAddress      string
	NetInterface   string
	Port           string
	ReadTimeout    time.Duration
	RPCDialTimeout time.Duration
	WriteTimeout   time.Duration
}

// Server represents API server instance
type Server struct {
	cfg               Config
	routeHandlers     map[string]http.HandlerFunc
	operationHandlers map[operation]operationHandler
	dlServiceClient   pb.DlClient
}

// NewServer returns new API server instance
func NewServer(cfg Config) *Server {
	server := Server{
		cfg: cfg,
	}
	server.routeHandlers = server.getRouteHandlers()
	server.operationHandlers = server.getOperationHandlers()
	return &server
}

// Serve starts GameDBV API server
func (s *Server) Serve(out io.Writer) error {
	closeDialConn, err := s.dialDlGrpc()
	if err != nil {
		return fmt.Errorf("could not dial Dl RPC: %w", err)
	}
	defer closeDialConn()

	router := s.initRouter()

	address := s.addressForServe()
	fmt.Fprintf(out, "API server address: %s\n", address) // TODO: implement logger instead of Fprintf in the wild..

	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: s.cfg.WriteTimeout,
		ReadTimeout:  s.cfg.ReadTimeout,
	}

	return srv.ListenAndServe()
}

func (s *Server) dialDlGrpc() (func() error, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	timeoutContext, cancel := context.WithTimeout(context.Background(), s.cfg.RPCDialTimeout)
	defer cancel()

	dlcRPCTarget := fmt.Sprintf("%s:%s", s.cfg.DlRPCAddress, s.cfg.DlRPCPort)
	conn, err := grpc.DialContext(timeoutContext, dlcRPCTarget, opts...) // this needs to be changed; maybe starting own process from api
	if err != nil {
		return nil, fmt.Errorf("grpc dial failure: %w", err)
	}

	s.dlServiceClient = pb.NewDlClient(conn)

	return conn.Close, nil
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
