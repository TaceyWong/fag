package cmd

import (
	"github.com/urfave/cli/v2"
)

// LatestCMD 获取最后一条的详情
var LatestCMD = &cli.Command{
	Name:    "latest",
	Aliases: []string{"ll"},
	Usage:   "获取最后一条的详情",
}
