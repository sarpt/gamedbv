package idx

import (
	"io"
	"log"

	"github.com/sarpt/gamedbv/pkg/platform"
)

const serverLoggerPrefix = "idx.Server#"

// Config instructs Idx how to download source file and where to put it.
type Config struct {
	Address   string
	ErrWriter io.Writer
	OutWriter io.Writer
	Port      string
	Indexes   map[platform.Variant]IndexConfig
}

type Server struct {
	cfg    Config
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
