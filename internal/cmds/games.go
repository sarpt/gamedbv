package cmds

import (
	"encoding/json"
	"fmt"

	"github.com/sarpt/gamedbv/pkg/db/queries"
)

const (
	gamesName = "gamedbv-games"
	// GamesResultStep is used for specifying that status includes information about results of game searching
	GamesResultStep = "games-results"
)

var (
	errOutputNotValidJSON = fmt.Errorf("games results output is not a valid json")
)

// Games is used to execute Games component binary
type Games struct {
	command command
}

// GamesCfg is used to control the behavior of command executing the Games component binary
type GamesCfg struct {
	Path string
}

// GamesArguments is used to provide arguments for command executing the Games component binary
type GamesArguments struct {
	Platforms []string
	Regions   []string
	Text      string
	Page      int
	PageLimit int
}

// GamesResultStatus is a result output from running games component
type GamesResultStatus struct {
	Step string              `json:"step"`
	Data queries.GamesResult `json:"data"`
}

// NewGames returns new Games cmd
func NewGames(cfg GamesCfg, args GamesArguments) Games {
	allArgs := parseGamesArguments(args)
	cmd := newCommand(gamesName, cfg.Path, allArgs)
	return Games{
		command: cmd,
	}
}

// Execute runs the command and waits for it to finish
func (games Games) Execute() (queries.GamesResult, error) {
	status := GamesResultStatus{}
	out, err := games.command.Stdout()
	if err != nil {
		return status.Data, err
	}

	if !json.Valid(out) {
		return status.Data, errOutputNotValidJSON
	}

	err = json.Unmarshal(out, &status)
	if err != nil {
		return status.Data, err
	}

	if status.Step != GamesResultStep {
		return status.Data, nil
	}

	return status.Data, nil
}

func parseGamesArguments(args GamesArguments) []string {
	allArgs := createJSONArgument(true)
	allArgs = append(allArgs, createTextArgument(args.Text)...)
	allArgs = append(allArgs, createPlatformsArguments(args.Platforms)...)
	allArgs = append(allArgs, createRegionsArguments(args.Regions)...)
	allArgs = append(allArgs, createPageArgument(args.Page)...)
	allArgs = append(allArgs, createLimitArgument(args.PageLimit)...)

	return allArgs
}
