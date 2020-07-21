package main

import (
	"flag"
	"fmt"
	"net"
	"sync"

	"github.com/sarpt/goutils/pkg/listflag"
	"google.golang.org/grpc"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/cmds"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/dl"
	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
	pb "github.com/sarpt/gamedbv/pkg/rpc/dl"
)

var jsonFlag *bool
var platformFlags *listflag.StringList
var grpcFlag *bool

const (
	defaultIP   = "localhost"
	defaultPort = "3004"
)

func init() {
	platformFlags = listflag.NewStringList([]string{})

	flag.Var(platformFlags, cmds.PlatformFlag, "platform specifies which console platform's database should be fetched")
	jsonFlag = flag.Bool(cmds.JSONFlag, false, "when specified as true, each line of output is presented as a json object")
	grpcFlag = flag.Bool(cmds.GRPCFlag, false, "when specified as true, the program launches in server mode, accepting gRPC requests and responding with streams of download process statuses")
	flag.Parse()
}

func main() {
	appCfg, err := config.NewApp()
	if err != nil {
		panic(err)
	}

	platformsToDownload, err := platformVariants(platformFlags.Values())
	if err != nil {
		panic(err)
	}

	if *grpcFlag {
		err = serveGRPC(appCfg)
	} else {
		executeOnce(appCfg, platformsToDownload)
	}

	if err != nil {
		panic(err)
	}
}

func serveGRPC(appCfg config.App) error {
	address := fmt.Sprintf("%s:%s", defaultIP, defaultPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterDlServer(grpcServer, newGRPCDlServer(appCfg))
	grpcServer.Serve(lis)

	return nil
}

func executeOnce(appCfg config.App, platforms []platform.Variant) {
	var printer progress.Notifier
	if *jsonFlag {
		printer = cli.NewJSONPrinter()
	} else {
		printer = cli.NewTextPrinter()
	}

	var wg sync.WaitGroup
	downloadPlatformSources(&wg, appCfg, platforms, printer)
	wg.Wait()
}

func downloadPlatformSources(wg *sync.WaitGroup, appCfg config.App, platforms []platform.Variant, notifier progress.Notifier) {
	for _, platformToDownload := range platforms {
		wg.Add(1)

		go func(platform platform.Variant) {
			defer wg.Done()
			dl.PlatformSource(appCfg.Dl(platform), platform, notifier)
		}(platformToDownload)
	}
}

func platformVariants(platforms []string) ([]platform.Variant, error) {
	var platformsToDownload []platform.Variant

	if len(platforms) == 0 {
		platformsToDownload = append(platformsToDownload, platform.All()...)
	} else {
		for _, val := range platforms {
			variant, err := platform.Get(val)
			if err != nil {
				return platformsToDownload, err
			}

			platformsToDownload = append(platformsToDownload, variant)
		}
	}

	return platformsToDownload, nil
}
