package commands

import (
	"fmt"

	"github.com/chroju/tfmodule/tfmodule"
	"github.com/mitchellh/cli"
)

// TemplateCommand
type TemplateCommand struct {
	UI cli.Ui
}

// Run runs template sub-command
func (c *TemplateCommand) Run(args []string) int {
	if len(args) != 1 {
		c.UI.Error(fmt.Sprintf("You must specify the module path.\n\n%s", helpTemplate))
		return 1
	}
	source := args[0]

	parser := tfmodule.NewParser(source)
	module, err := parser.ParseTfModule(source)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Output(module.String())

	return 0
}

func (c *TemplateCommand) Help() string {
	return helpTemplate
}

func (c *TemplateCommand) Synopsis() string {
	return "Parse Terraform module files and output module template."
}

const helpTemplate = `
Usage: tfmodule template [options] SOURCE

  Output the Terraform module template with given module source path.

Options:

  --minimum    Ouptput template does not include the variables which has a default value.
`
