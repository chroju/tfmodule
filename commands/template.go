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
	var isNoDefault bool
	var isNoOutputs bool
	var isMinimum bool
	buf := &bytes.Buffer{}
	f := flag.NewFlagSet("template", flag.ContinueOnError)
	f.SetOutput(buf)
	f.StringVarP(&name, "name", "n", "", "module name")
	f.BoolVar(&isNoDefault, "no-defaults", false, "print template without variables with default values")
	f.BoolVar(&isNoOutputs, "no-outputs", false, "print template without outputs")
	f.BoolVar(&isMinimum, "minimum", false, "print minimum template (same as --no-outputs and --no-defaults)")
	if err := f.Parse(flagArgs); err != nil {
		c.UI.Error(helpTemplate)
		return 1
	}
	if isMinimum {
		isNoDefault = true
		isNoOutputs = true
	}

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
	if name != "" {
		module.Name = name
	}
	c.UI.Output(module.PrintModuleTemplate(isNoDefault, isNoOutputs))

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
  --no-defaults            Print template without variables with default values
  --no-outputs             Print template without outputs
  --minimum                Print minimum template (same as --no-outputs and --no-defaults)
`
