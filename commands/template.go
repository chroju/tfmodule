package commands

import (
	"fmt"

	"github.com/chroju/tfmodule/tfmodule"
	"github.com/mitchellh/cli"
	flag "github.com/spf13/pflag"
)

// TemplateCommand
type TemplateCommand struct {
	UI           cli.Ui
	Name         string
	IsNoDefaults bool
	IsNoOutputs  bool
	IsMinimum    bool
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
	f := flag.NewFlagSet("template", flag.ContinueOnError)
	f.StringVarP(&c.Name, "name", "n", "", "module name")
	f.BoolVar(&c.IsNoDefaults, "no-defaults", false, "print template without variables with default values")
	f.BoolVar(&c.IsNoOutputs, "no-outputs", false, "print template without outputs")
	f.BoolVar(&c.IsMinimum, "minimum", false, "print minimum template (same as --no-outputs and --no-defaults)")
	if err := f.Parse(flagArgs); err != nil {
		c.UI.Error(helpTemplate)
		return 1
	}
	if c.IsMinimum {
		c.IsNoDefaults = true
		c.IsNoOutputs = true
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

	if c.Name != "" {
		module.Name = c.Name
	}

	options := &tfmodule.PrintOptions{
		Format:       "template",
		IsNoDefaults: c.IsNoDefaults,
		IsNoOutputs:  c.IsNoOutputs,
	}

	printer, err := tfmodule.NewPrinter(module, options)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	out, err := printer.Print()
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	c.UI.Output(out)

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
