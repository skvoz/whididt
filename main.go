package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/joho/godotenv"
	lib "github.com/skvoz/whididt/lib"
	cli "github.com/urfave/cli/v2"
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

func init() {
	var err error
	if os.Getenv("APP_ENV") == "prod" {
		err = godotenv.Load(".prod.env")
	} else {
		err = godotenv.Load(".env")
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
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
			var path = []string{}
			if len(cCtx.StringSlice("path")) > 0 {
				path = cCtx.StringSlice("path")
			} else {
				path = []string{lib.GetProjectPath()}
			}

			tmpl, err := template.ParseFiles(filepath.Join("../",
				"whididt.html"))
			if err != nil {
				panic(err)
			}
			td := TemplateData{
				BossName:   "Fedia",
				ReportDate: "11-11-2023",
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
