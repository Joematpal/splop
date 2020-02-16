package flags

import "github.com/urfave/cli/v2"

const FilePath = "file"

var FilePathFlag = &cli.StringFlag{
	Name: FilePath,
}

const Url = "url"

var UrlFlag = &cli.StringFlag{
	Name: Url,
}
