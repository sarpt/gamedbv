package idx

import (
	"log"

	"github.com/sarpt/gamedbv/internal/progress"
	pb "github.com/sarpt/gamedbv/pkg/rpc/idx"
)

type initializeDatabaseNotifier struct {
	errLog *log.Logger
	stream pb.Idx_InitializeDatabaseServer
	outLog *log.Logger
}

// NextStatus sends information about new status.
func (n initializeDatabaseNotifier) NextStatus(status progress.Status) {
	res := pb.InitializeDatabaseStatus{
		Step:    status.Step,
		Message: status.Message,
	}

	err := n.stream.Send(&res)
	if err != nil {
		n.errLog.Printf("could not send through grpc: %v\n", err)
	}
}

// NextError to be implemented
func (n initializeDatabaseNotifier) NextError(err error) {}

// preparePlatformNotifier is used to send information through gRPC stream.
type preparePlatformNotifier struct {
	errLog *log.Logger
	stream pb.Idx_PreparePlatformsServer
	outLog *log.Logger
}

// NextStatus sends information about new status.
func (n preparePlatformNotifier) NextStatus(status progress.Status) {
	res := pb.PreparePlatformsStatus{
		Platform: status.Platform,
		Process:  status.Process,
		Step:     status.Step,
		Message:  status.Message,
	}

	err := n.stream.Send(&res)
	if err != nil {
		n.errLog.Printf("could not send through grpc: %v\n", err)
	}
}

// NextError to be implemented
func (n preparePlatformNotifier) NextError(err error) {}
