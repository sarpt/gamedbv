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
	preparePlatform func(platform.Variant, progress.Notifier) error
	errLog          *log.Logger
	outLog          *log.Logger
}

func newGRPCServer(cfg gRPCConfig) *gRPCServer {
	return &gRPCServer{
		preparePlatform: cfg.preparePlatform,
		errLog:          cfg.errLog,
		outLog:          cfg.outLog,
	}
}

// DownloadPlatforms handles gRPC request to download one or multiple platforms.
func (s *gRPCServer) PreparePlatforms(req *pb.PreparePlatformsReq, stream pb.Idx_PreparePlatformsServer) error {
	s.outLog.Printf("incoming gRPC request for PreparePlatforms: %s\n", strings.Join(req.GetPlatforms(), ", "))

	platforms, err := platform.ByNames(req.GetPlatforms())
	if err != nil {
		return err
	}

	nCfg := gRPCNotifierConfig{
		errLog: s.errLog,
		stream: stream,
		outLog: s.outLog,
	}
	notifier := newGRPCNotifier(nCfg)

	var wg sync.WaitGroup
	for _, platformToDownload := range platforms {
		wg.Add(1)

		go func(platform platform.Variant) {
			defer wg.Done()
			err = s.preparePlatform(platform, notifier)
			if err != nil {
				s.errLog.Printf("could not prepare platform: %v", err)
			}
		}(platformToDownload)
	}
	wg.Wait()

	return nil
}
