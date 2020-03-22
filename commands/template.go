package commands

import (
	"bytes"
	"fmt"

	"github.com/chroju/tfmodule/tfmodule"
	"github.com/mitchellh/cli"
	flag "github.com/spf13/pflag"
)

// TemplateCommand
type TemplateCommand struct {
	UI cli.Ui
}

// Run runs template sub-command
func (c *TemplateCommand) Run(args []string) int {
	if len(args) == 0 {
		c.UI.Error(fmt.Sprintf("You must specify the module path.\n\n%s", helpTemplate))
		return 1
	}
	source := args[0]
	flagArgs := args[1:]

	// flags
	var name string
	buf := &bytes.Buffer{}
	f := flag.NewFlagSet("template", flag.ContinueOnError)
	f.SetOutput(buf)
	f.StringVarP(&name, "name", "n", "", "module name")
	if err := f.Parse(flagArgs); err != nil {
		c.UI.Error(helpTemplate)
		return 1
	}

	parser := tfmodule.NewParser()
	module, err := parser.ParseTfModule(source)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	if name != "" {
		module.Name = name
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
Usage: tfmodule template SOURCE [options]

  Output the Terraform module template with given module source path.

Options:

  --name=modulename, -n    Name of module
  --minimum                Ouptput template does not include the variables which has a default value.
`
