package cmd

import (
	"github.com/joematpal/splop/cmd/flags"
	cli "github.com/urfave/cli/v2"
)

// NewApp is the constuctor of a houshound cli
func NewApp() *cli.App {
	flgs := []cli.Flag{
		flags.FilePathFlag,
		flags.UrlFlag,
	}

	commands := []*cli.Command{
		addCmd,
	}

	return &cli.App{
		Name:     "splop",
		Usage:    "",
		Flags:    flgs,
		Commands: commands,
	}
}
