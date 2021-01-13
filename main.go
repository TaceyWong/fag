package main

import (
	"log"
	"os"
	"sort"

	"github.com/TaceyWong/fag/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "fag"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "王新勇(Tacey Wong)",
			Email: "xinyong.wang@qq.com",
		},
	}
	app.Copyright = " © 2020 - 而今 " + app.Authors[0].String()
	app.Usage = "关于Golang信息的聚合"
	app.Description = `目前聚合列表如下:
	+ awesome-go.com
	+ golangweekly.com/issues
	`
	app.EnableBashCompletion = true
	app.HideHelpCommand = true

	app.Commands = []*cli.Command{
		cmd.LatestCMD,
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	sort.Sort(cli.FlagsByName(app.Flags))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
