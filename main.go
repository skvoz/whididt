package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	lib "github.com/skvoz/whididt/lib"
	cli "github.com/urfave/cli/v2"

	"time"
)

type TemplateData struct {
	BossName   string
	ReportDate string
	PD         []ProjectData
}

type ProjectData struct {
	ProjectName string
	CommitData  []string
}

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
			currentTime := time.Now()

			formattedDate := currentTime.Format("01-02-2006")

			var path = []string{}
			var bossName string = ""
			var reportDateStart string = formattedDate
			var reportDateUntil string = ""

			if len(cCtx.StringSlice("path")) > 0 {
				path = cCtx.StringSlice("path")
			} else {
				path = []string{lib.GetProjectPath()}
			}

			if cCtx.String("boss") != "" {
				bossName = cCtx.String("boss")
			}

			if cCtx.String("start") != "" {
				reportDateStart = cCtx.String("start")
			}

			if cCtx.String("until") != "" {
				reportDateUntil = cCtx.String("until")
			}

			tmpl, err := template.ParseFiles(filepath.Join("templates",
				"base.html"))

			if err != nil {
				panic(err)
			}

			td := TemplateData{
				BossName:   bossName,
				ReportDate: reportDateStart + " " + reportDateUntil,
			}

			for _, v := range path {
				gitCommit, _ := lib.GetGitCommitData(v, start, until)
				projectName, _ := lib.GetGitProjectName(v)
				td.PD = append(td.PD, ProjectData{projectName, gitCommit})
			}

			err = tmpl.Execute(os.Stdout, td)
			if err != nil {
				panic(err)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
