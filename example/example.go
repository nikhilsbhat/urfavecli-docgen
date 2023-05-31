package main

import (
	"fmt"
	"log"

	docgen "github.com/nikhilsbhat/urfavecli-docgen"
	"github.com/urfave/cli/v2"
)

func main() {
	appCli := cli.App{
		Name:                 "sample-app",
		Usage:                "Utility to run some random stuff",
		UsageText:            "sample-app [flags]",
		EnableBashCompletion: true,
		HideHelp:             false,
		Authors: []*cli.Author{
			{
				Name:  "Nikhil Bhat",
				Email: "nikhilsbhat93@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "version of the sample-app",
				Action: func(context *cli.Context) error {
					fmt.Println("v0.0.2")

					return nil
				},
			},
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "log-level",
				Usage:   "set log level for the sample app",
				Aliases: []string{"log"},
				Value:   "info",
			},
		},
		Action: func(context *cli.Context) error {
			fmt.Println("from sample app")

			return nil
		},
	}

	if err := docgen.GenerateDocs(&appCli, "sample_app"); err != nil {
		log.Fatalln(err)
	}
}
