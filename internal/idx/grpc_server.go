package idx

import (
	"log"
	"strings"
	"sync"

	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
	pb "github.com/sarpt/gamedbv/pkg/rpc/idx"
)

type gRPCConfig struct {
	preparePlatform func(platform.Variant, progress.Notifier) error
	errLog          *log.Logger
	outLog          *log.Logger
}

// gRPCServer provides mechanisms to use Dl service both through RPC and by commands.
type gRPCServer struct {
	pb.UnimplementedIdxServer
	preparePlatform    func(platform.Variant, progress.Notifier) error
	initializeDatabase func(string, string, bool, progress.Notifier) error
	errLog             *log.Logger
	outLog             *log.Logger
}

func newGRPCServer(cfg gRPCConfig) *gRPCServer {
	return &gRPCServer{
		preparePlatform: cfg.preparePlatform,
		errLog:          cfg.errLog,
		outLog:          cfg.outLog,
	}
}

func (s *gRPCServer) InitializeDatabase(req *pb.InitializeDatabaseReq, stream pb.Idx_InitializeDatabaseServer) error {
	s.outLog.Printf("incoming gRPC request for InitializeDatabase with variant '%s' in path '%s' and with force set to %t\n", req.GetVariant(), req.GetPath(), req.GetForce())

	notifier := initializeDatabaseNotifier{
		errLog: s.errLog,
		stream: stream,
		outLog: s.outLog,
	}

	err := s.initializeDatabase(req.GetPath(), req.GetVariant(), req.GetForce(), notifier)
	if err != nil {
		s.errLog.Printf("could not initialize database: %v", err)
	}

	return err
}

// DownloadPlatforms handles gRPC request to download one or multiple platforms.
func (s *gRPCServer) PreparePlatforms(req *pb.PreparePlatformsReq, stream pb.Idx_PreparePlatformsServer) error {
	s.outLog.Printf("incoming gRPC request for PreparePlatforms: %s\n", strings.Join(req.GetPlatforms(), ", "))

	platforms, err := platform.ByNames(req.GetPlatforms())
	if err != nil {
		return err
	}

	notifier := preparePlatformNotifier{
		errLog: s.errLog,
		stream: stream,
		outLog: s.outLog,
	}

	var wg sync.WaitGroup
	for _, platformToDownload := range platforms {
		wg.Add(1)

		go func(platform platform.Variant) {
			defer wg.Done()
			err = s.preparePlatform(platform, notifier) // aggregate errors
			if err != nil {
				s.errLog.Printf("could not prepare platform: %v", err)
			}
		}(platformToDownload)
	}
	wg.Wait()

	return nil
}
