package tfmodule

import (
	"errors"
)

// Printer is an interface to print terraform module.
type Printer interface {
	Print() (string, error)
}

type PrintOptions struct {
	Format       string
	IsNoDefaults bool
	IsNoOutputs  bool
}

func NewPrinter(m *Module, o *PrintOptions) (Printer, error) {
	var p Printer

	switch o.Format {
	case "template":
		p = newTemplatePrinter(m, o.IsNoDefaults, o.IsNoOutputs)
	case "analysis":
		p = newAnalysisPrinter(m)
	default:
		return nil, errors.New("Invalid printer format.")
	}

	return p, nil
}
