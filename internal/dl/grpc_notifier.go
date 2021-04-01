package dl

import (
	"log"

	"github.com/sarpt/gamedbv/internal/progress"
	pb "github.com/sarpt/gamedbv/pkg/rpc/dl"
)

// gRPCNotifier is used to send information through gRPC stream
type gRPCNotifier struct {
	errLog *log.Logger
	stream pb.Dl_DownloadPlatformsServer
	outLog *log.Logger
}

// NextStatus sends information about new status
func (n gRPCNotifier) NextStatus(status progress.Status) {
	res := pb.PlatformsDownloadStatus{
		Platform: status.Platform,
		Process:  status.Process,
		Step:     status.Step,
		Message:  status.Message,
	}

	err := n.stream.Send(&res)
	if err != nil {
		n.errLog.Printf("could not send error through grpc: %v\n", err)
	}
}

// NextError to be implemented
func (n gRPCNotifier) NextError(err error) {}
