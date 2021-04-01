package dl

import (
	"log"
	"strings"
	"sync"

	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
	pb "github.com/sarpt/gamedbv/pkg/rpc/dl"
)

type gRPCConfig struct {
	downloadPlatform func(platform.Variant, progress.Notifier) error
	errLog           *log.Logger
	outLog           *log.Logger
}

// gRPCServer provides mechanisms to use Dl service both through RPC and by commands.
type gRPCServer struct {
	pb.UnimplementedDlServer
	downloadPlatform func(platform.Variant, progress.Notifier) error
	errLog           *log.Logger
	outLog           *log.Logger
}

func newGRPCServer(cfg gRPCConfig) *gRPCServer {
	return &gRPCServer{
		downloadPlatform: cfg.downloadPlatform,
		errLog:           cfg.errLog,
		outLog:           cfg.outLog,
	}
}

// DownloadPlatforms handles gRPC request to download one or multiple platforms.
func (s *gRPCServer) DownloadPlatforms(req *pb.PlatformsDownloadReq, stream pb.Dl_DownloadPlatformsServer) error {
	s.outLog.Printf("incoming gRPC request for DownloadPlatforms: %s\n", strings.Join(req.GetPlatforms(), ", "))

	platforms, err := platform.ByNames(req.GetPlatforms())
	if err != nil {
		return err
	}

	notifier := gRPCNotifier{
		errLog: s.errLog,
		stream: stream,
		outLog: s.outLog,
	}

	var wg sync.WaitGroup
	for _, platformToDownload := range platforms {
		wg.Add(1)

		go func(platform platform.Variant) {
			defer wg.Done()
			s.downloadPlatform(platform, notifier)
		}(platformToDownload)
	}
	wg.Wait()

	return nil
}
