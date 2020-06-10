package cmds

import "io"

const (
	gamesName = "gamedbv-games"
)

// Games is used to execute Games component binary
type Games struct {
	command command
}

// GamesCfg is used to control the behavior of command executing the Games component binary
type GamesCfg struct {
	Path      string
	Output    io.Writer
	ErrOutput io.Writer
}

// GamesArguments is used to provide arguments for command executing the Games component binary
type GamesArguments struct {
	Platforms []string
	Regions   []string
	TextQuery string
	page      int
	pageLimit int
}

// NewGames returns new Games cmd
func NewGames(cfg GamesCfg, args GamesArguments) Games {
	allArgs := createJSONArgument(true)
	allArgs = append(allArgs, createPlatformsArguments(args.Platforms)...)

	cmd := newCommand(idxName, cfg.Path, allArgs, cfg.Output, cfg.ErrOutput)
	return Games{
		command: cmd,
	}
}

// Execute runs the command and waits for it to finish
func (games Games) Execute() error {
	return games.command.Execute()
}
