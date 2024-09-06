// Package cmd contains main functionality for the application using CLI commands.
package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

// Run executes the application.
func Run() error {
	app := cli.App{
		Name:      "blitzbot",
		Usage:     "Wassup my telegram bot!",
		UsageText: "blitzbot [command] serve (s) [command option] debug (d)",
		Version:   "0.1",
		Commands: []*cli.Command{
			serveCMD,
		},
	}

	return app.Run(os.Args)
}
