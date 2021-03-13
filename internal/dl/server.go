package dl

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
	pb "github.com/sarpt/gamedbv/pkg/rpc/dl"
	"google.golang.org/grpc"
)

const (
	serverLoggerPrefix = "dl.Server#"
)

// FilesStatus groups information about existence of specific platform's source files.
type FilesStatus struct {
	SourceExists bool
}

// Config instructs Dl how to download source file and where to put it.
type Config struct {
	Address   string
	ErrWriter io.Writer
	OutWriter io.Writer
	Port      string
	Sources   map[platform.Variant]SourceConfig
}

type Server struct {
	cfg    Config
	errLog *log.Logger
	outLog *log.Logger
}

type SourceConfig struct {
	DirectoryPath   string
	Filepath        string
	ForceRedownload bool
	PlatformName    string
	URL             string
}

func NewServer(cfg Config) *Server {
	return &Server{
		cfg:    cfg,
		errLog: log.New(cfg.ErrWriter, serverLoggerPrefix, log.LstdFlags),
		outLog: log.New(cfg.OutWriter, serverLoggerPrefix, log.LstdFlags),
	}
}

// DownloadPlatformSource downloads neccessary source files related to provided platform.
// Progress of the process is being reported on the printer.
func (s *Server) DownloadPlatformSource(variant platform.Variant, printer progress.Notifier) {
	sourceCfg := s.cfg.Sources[variant]
	sourcesFilesStatuses, err := getFilesStatuses(sourceCfg)
	if err != nil {
		printer.NextError(err)
		return
	}

	if sourcesFilesStatuses.SourceExists && !sourceCfg.ForceRedownload {
		printer.NextStatus(newArchiveFileAlreadyPresentStatus(variant))
		return
	}

	err = preparePlatformDirectory(sourceCfg)
	if err != nil {
		printer.NextError(err)
		return
	}

	printer.NextStatus(newDownloadingInProgressStatus(variant))
	err = downloadSourceFile(sourceCfg)
	if err != nil {
		printer.NextError(err)
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
		downloadPlatform: s.DownloadPlatformSource,
		outLog:           s.outLog,
		errLog:           s.errLog,
	}
	pb.RegisterDlServer(server, newGRPCServer(grpcServerCfg))

	s.outLog.Printf("gRPC server listening on %s\n", address)
	server.Serve(lis)

	return nil
}

func downloadSourceFile(cfg SourceConfig) error {
	filePath := cfg.Filepath
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = getFile(cfg.URL, outputFile)
	return err
}

func getFilesStatuses(cfg SourceConfig) (FilesStatus, error) {
	var filesStatus FilesStatus

	filePath := cfg.Filepath
	filesStatus.SourceExists = fileExists(filePath)

	return filesStatus, nil
}

func preparePlatformDirectory(cfg SourceConfig) error {
	directory := cfg.DirectoryPath

	err := os.MkdirAll(directory, 0700)
	return err
}
