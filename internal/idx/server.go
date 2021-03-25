package idx

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/sarpt/gamedbv/pkg/db"
	"github.com/sarpt/gamedbv/pkg/platform"
	pb "github.com/sarpt/gamedbv/pkg/rpc/idx"
	"google.golang.org/grpc"
)

const serverLoggerPrefix = "idx.Server#"

// Config instructs Idx how to download source file and where to put it.
type Config struct {
	Address    string
	DbPath     string
	DbVariant  string
	DbMaxLimit int
	ErrWriter  io.Writer
	OutWriter  io.Writer
	Port       string
	Indexes    map[platform.Variant]IndexConfig
}

type Server struct {
	cfg    Config
	db     *db.Database
	errLog *log.Logger
	outLog *log.Logger
}

func NewServer(cfg Config) *Server {
	return &Server{
		cfg:    cfg,
		errLog: log.New(cfg.ErrWriter, serverLoggerPrefix, log.LstdFlags),
		outLog: log.New(cfg.OutWriter, serverLoggerPrefix, log.LstdFlags),
	}
}

func (s *Server) ServeGRPC() error {
	address := fmt.Sprintf("%s:%s", s.cfg.Address, s.cfg.Port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	grpcServerCfg := gRPCConfig{
		preparePlatform: s.PreparePlatform,
		outLog:          s.outLog,
		errLog:          s.errLog,
	}
	pb.RegisterIdxServer(server, newGRPCServer(grpcServerCfg))

	s.outLog.Printf("gRPC server listening on %s\n", address)
	server.Serve(lis)

	return nil
}
