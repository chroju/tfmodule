package main

import (
	"fmt"
	"os"

	"github.com/chroju/tfmodule/commands"
	"github.com/mitchellh/cli"
)

const (
	app     = "tfmodule"
	version = "0.1.0"
)

func main() {
	c := cli.NewCLI(app, version)
	c.Args = os.Args[1:]
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c.Commands = map[string]cli.CommandFactory{
		"template": func() (cli.Command, error) {
			return &commands.TemplateCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
		"analyze": func() (cli.Command, error) {
			return &commands.AnalyzeCommand{UI: &cli.ColoredUi{Ui: ui, ErrorColor: cli.UiColorRed}}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		ui.Error(fmt.Sprintf("Error: %s", err))
	}

	os.Exit(exitStatus)
}
