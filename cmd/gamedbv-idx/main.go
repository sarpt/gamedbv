package main

import (
	"flag"
	"os"
	"sync"

	"github.com/sarpt/goutils/pkg/listflag"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/cmds"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/idx"
	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/platform"
)

var grpcFlag *bool
var jsonFlag *bool
var platformFlags *listflag.StringList

func init() {
	platformFlags = listflag.NewStringList([]string{})

	flag.Var(platformFlags, cmds.PlatformFlag, "platform specifies which console platform's database should be fetched")
	jsonFlag = flag.Bool(cmds.JSONFlag, false, "when specified as true, each line of output is presented as a json object")
	grpcFlag = flag.Bool(cmds.GRPCFlag, false, "when specified as true, the program launches in server mode, accepting gRPC requests and responding with streams of download process statuses")
	flag.Parse()
}

func main() {
	projectCfg, err := config.Create()
	if err != nil {
		panic(err)
	}

	cfg := projectCfg.Idx
	cfg.ErrWriter = os.Stderr
	cfg.OutWriter = os.Stdout
	server := idx.NewServer(cfg)

	err = server.OpenDatabase()
	if err != nil {
		panic(err)
	}

	if *grpcFlag {
		err = server.ServeGRPC()
	} else {
		var platformsToParse []platform.Variant

		var printer progress.Notifier
		if *jsonFlag {
			printer = cli.NewJSONPrinter()
		} else {
			printer = cli.NewTextPrinter()
		}

		if len(platformFlags.Values()) == 0 {
			platformsToParse = append(platformsToParse, platform.All()...)
		} else {
			for _, val := range platformFlags.Values() {
				variant, err := platform.Get(val)
				if err != nil {
					panic(err)
				}

				platformsToParse = append(platformsToParse, variant)
			}
		}

		var wg sync.WaitGroup
		for _, platformToParse := range platformsToParse {
			wg.Add(1)

			go func(platform platform.Variant) {
				defer wg.Done()
				server.PreparePlatform(platform, printer)
			}(platformToParse)
		}

		wg.Wait()
	}

	if err != nil {
		panic(err)
	}
}
