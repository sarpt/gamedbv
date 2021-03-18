package main

import (
	"flag"

	"github.com/sarpt/gamedbv/internal/cli"
	"github.com/sarpt/gamedbv/internal/cmds"
	"github.com/sarpt/gamedbv/internal/config"
	"github.com/sarpt/gamedbv/internal/progress"
	"github.com/sarpt/gamedbv/pkg/db"
)

var (
	listFlag *string
	jsonFlag *bool
)

const (
	statusResultStep = "status-result"
	platformsStatus  = "platforms"
	regionsStatus    = "regions"
	languagesStatus  = "languages"
)

func init() {
	listFlag = flag.String(cmds.ListFlag, "", "specify what type of list should be shown")
	jsonFlag = flag.Bool(cmds.JSONFlag, false, "show output as a json")
	flag.Parse()
}

func main() {
	var printer progress.Notifier
	if *jsonFlag {
		printer = cli.NewJSONPrinter()
	} else {
		printer = cli.NewTextPrinter()
	}

	cfg, err := config.Create()
	if err != nil {
		panic(err)
	}

	var status progress.Status
	switch *listFlag {
	case platformsStatus:
		status, err = handlePlatforms(cfg.Database)
	case regionsStatus:
		status, err = handleRegions(cfg.Database)
	case languagesStatus:
		status, err = handleLanguages(cfg.Database)
	default:
		flag.PrintDefaults()
		return
	}

	if err != nil {
		panic(err)
	}
	printer.NextStatus(status)
}

func handleRegions(dbCfg db.Config) (progress.Status, error) {
	return progress.Status{}, nil // TODO: needs implementation
}

func handleLanguages(dbCfg db.Config) (progress.Status, error) {
	return progress.Status{}, nil // TODO: needs implementation
}

func prepareResultStatus(data interface{}, out string) progress.Status {
	if *jsonFlag {
		return progress.Status{
			Step: statusResultStep,
			Data: data,
		}
	}

	return progress.Status{
		Message: out,
	}
}
