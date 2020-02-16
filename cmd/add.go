package cmd

import (
	"github.com/joematpal/splop/cmd/flags"
	"github.com/joematpal/splop/lib/geojson"
	cli "github.com/urfave/cli/v2"
)

var addCmd = &cli.Command{
	Name: "add",
	// Flags: []cli.Flag{},
	Action: func(ctx *cli.Context) error {
		url := ctx.String(flags.Url)
		filePath := ctx.String(flags.FilePath)
		return geojson.LoadGeoJson(url, filePath)
	},
}
