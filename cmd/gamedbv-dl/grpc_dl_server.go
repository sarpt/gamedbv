package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/progress"
	pb "github.com/sarpt/gamedbv/pkg/rpc/dl"
)

// GRPCDlServer implements server side of Dl service
type GRPCDlServer struct {
	pb.UnimplementedDlServer
	appCfg config.App
}

func newGRPCDlServer(appCfg config.App) *GRPCDlServer {
	return &GRPCDlServer{
		appCfg: appCfg,
	}
}

// DownloadPlatforms handles request to download one or multiple platforms
func (s *GRPCDlServer) DownloadPlatforms(req *pb.PlatformsDownloadReq, stream pb.Dl_DownloadPlatformsServer) error {
	platforms, err := platformVariants(req.GetPlatforms())
	if err != nil {
		return err
	}

	notifier := newGRPCNotifier(stream)

	var wg sync.WaitGroup
	downloadPlatformSources(&wg, s.appCfg, platforms, notifier)
	wg.Wait()

	return nil
}

func newGRPCNotifier(stream pb.Dl_DownloadPlatformsServer) GRPCNotifier {
	return GRPCNotifier{
		stream: stream,
	}
}

// GRPCNotifier is used to send information through gRPC stream
type GRPCNotifier struct {
	stream pb.Dl_DownloadPlatformsServer
}

// NextStatus sends information about new status
func (n GRPCNotifier) NextStatus(status progress.Status) {
	platformDownloadStatus := pb.PlatformsDownloadStatus{
		Platform: status.Platform,
		Process:  status.Process,
		Step:     status.Step,
		Message:  status.Message,
	}

	err := n.stream.Send(&platformDownloadStatus)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not send error through grpc: %v", err)
	}
}

// NextError to be implemented
func (n GRPCNotifier) NextError(err error) {}
