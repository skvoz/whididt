package main

import (
	"fmt"
	"log"
	"os"
	"time"

	lib "github.com/skvoz/whididt/lib"
	cli "github.com/urfave/cli/v2"
)

func main() {
	var start string
	var until string
	var boss string

	app := &cli.App{
		Name:  "whididt",
		Usage: "Generate daylick report.",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "path",
				Aliases: []string{"p"},
				Usage:   "`PATH` to your project. Can use multiple times.",
			},
			&cli.StringFlag{
				Name:        "start",
				Aliases:     []string{"S"},
				Usage:       "date start log, (ex: -S 01-01-2022) get log from 01-01-2022 to 02-01-2022",
				Destination: &start,
			},
			&cli.StringFlag{
				Name:        "until",
				Aliases:     []string{"u"},
				Usage:       "date until log,  (ex: -S 01-01-2022 -u 05-01-2022) get log from 01-01-2022 to 05-01-2022, if use -S , default unitl = one day",
				Destination: &until,
			},
			&cli.StringFlag{
				Name:        "boss",
				Value:       "",
				Aliases:     []string{"b"},
				Usage:       "`BOSS` name will output in report",
				Destination: &boss,
			},
		},
		Action: func(cCtx *cli.Context) error {
			var path = []string{}
			if len(cCtx.StringSlice("path")) > 0 {
				path = cCtx.StringSlice("path")
			} else {
				path = []string{lib.GetProjectPath()}
			}
			currentTime := time.Now()
			var templateHeader string
			var templateBody string

			templateHeader = `Hello, %s

Daylick report. Date: %s
`
			templateBody = `
Project: %s
`
			var header, body string

			header = fmt.Sprintf(templateHeader, boss, currentTime.Format("01-02-2006"))
			for _, v := range path {
				gitCommit, _ := lib.GetGitCommitData(v, start, until)
				projectName, _ := lib.GetGitProjectName(v)

				body += fmt.Sprintf(templateBody, projectName)
				for _, n := range gitCommit {
					if n != "" {
						body += fmt.Sprintf(" - %s\n", n)
					}
				}
			}

			fmt.Print(header, body)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
