package commands

import (
	"fmt"

	"github.com/chroju/tfmodule/tfmodule"
	"github.com/mitchellh/cli"
)

// AnalyzeCommand
type AnalyzeCommand struct {
	UI cli.Ui
}

// Run runs analyze sub-command
func (c *AnalyzeCommand) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error(fmt.Sprintf("You must specify the module path.\n\n%s", helpAnalyze))
		return 1
	}
	source := args[0]

	parser, err := tfmodule.NewParser(source)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	module, err := parser.Parse()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Output(module.PrintModuleAnalysis())

	return 0
}

func (c *AnalyzeCommand) Help() string {
	return helpAnalyze
}

func (c *AnalyzeCommand) Synopsis() string {
	return "Analyze the Terraform module and output its summary."
}

const helpAnalyze = `
Usage: tfmodule analyze SOURCE

  Analyze the Terraform module and output its summary.
`
