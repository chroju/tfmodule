package tfmodule

import (
	"sort"
	"strings"
)

type analysisPrinter struct {
	Module *Module
}

func newAnalysisPrinter(m *Module) *analysisPrinter {
	return &analysisPrinter{
		Module: m,
	}
}

func (p *analysisPrinter) Print() (string, error) {
	emptyLine := []string{""}
	results := []string{"resources:"}

	resources := make([]string, len(p.Module.Resources))
	for i, r := range p.Module.Resources {
		resources[i] = "  " + r.Type + "." + r.Name
	}
	sort.Strings(resources)
	results = append(results, resources...)
	results = append(results, emptyLine...)
	results = append(results, "outputs:")

	outputs := make([]string, len(p.Module.Outputs))
	for i, o := range p.Module.Outputs {
		outputs[i] = " " + string(o.Value.Bytes())
	}
	sort.Strings(outputs)
	results = append(results, outputs...)
	results = append(results, emptyLine...)

	return strings.Join(results, "\n"), nil
}
