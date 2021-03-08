package api

import (
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
	IPAddress    string
	Port         string
	NetInterface string
	Debug        bool
	GamesConfig  games.Config
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type handlerCreator func() http.HandlerFunc

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
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:3004", opts...) // this needs to be changed; need to add timeout for dial, some kind of config reading for dl-service process, maybe starting own process from api
	if err != nil {
		return fmt.Errorf("fail to dial: %v", err)
	}

	defer conn.Close()
	s.dlServiceClient = pb.NewDlClient(conn)

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
